import React from 'react';
import { Fragment, useRef, useState } from 'react'
import { Dialog, Transition } from '@headlessui/react'
import { ServerIcon, CheckIcon } from '@heroicons/react/outline'
import { connectServer } from '../services/apis/system.query';
import { setServer } from '../services/store/connectivity/server'

interface ConnectState {
    url: string;
    isConnected: boolean;
    error: string;
    sysInfo: any;
    initFocusRef: any;
    open: boolean;
    connecting: boolean;
}

class ConnectPage extends React.Component<any, ConnectState>{
    initFocus: any = null
    constructor(props: any) {
        super(props);
        props = {
            handler: props.handler,
        };
        this.state = {
            url: "http://127.0.0.1:17415",
            isConnected: false,
            error: "",
            sysInfo: {},
            initFocusRef: null,
            open: true,
            connecting: false,
        };
    }
    componentDidMount() {
        this.initFocus = React.createRef()
    }
    handleConnect = () => {
        let self = this
        this.setState({ connecting: true })
        connectServer(this.state.url + "/ping").then(function (res) {
            self.setState({ sysInfo: res.data })
            self.setState({ isConnected: true })
        })
    }

    handleUrlChange = (e: any) => {
        this.setState({ url: e.target.value })
    }
    handleCompleteConnection = () => {
        this.setState({ open: false })
        this.props.handler(setServer(this.state.url))
    }
    handleCancel = () => {
        if (this.state.isConnected) {
            this.setState({ isConnected: false })
            this.setState({ connecting: false })
        }
    }
    render() {
        return (
            <Transition.Root show={this.state.open} as={Fragment}>
                <Dialog
                    as="div"
                    static
                    className="fixed z-10 inset-0 overflow-y-auto"
                    initialFocus={this.initFocus}
                    open={this.state.open}
                    onClose={() => this.setState({ open: false })}
                >
                    <div className="flex items-end justify-center min-h-screen pt-4 px-4 pb-20 text-center sm:block sm:p-0">
                        <Transition.Child
                            as={Fragment}
                            enter="ease-out duration-300"
                            enterFrom="opacity-0"
                            enterTo="opacity-100"
                            leave="ease-in duration-200"
                            leaveFrom="opacity-100"
                            leaveTo="opacity-0"
                        >
                            <Dialog.Overlay className="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity" />
                        </Transition.Child>

                        {/* This element is to trick the browser into centering the modal contents. */}
                        <span className="hidden sm:inline-block sm:align-middle sm:h-screen" aria-hidden="true">
                            &#8203;
                        </span>
                        <Transition.Child
                            as={Fragment}
                            enter="ease-out duration-300"
                            enterFrom="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"
                            enterTo="opacity-100 translate-y-0 sm:scale-100"
                            leave="ease-in duration-200"
                            leaveFrom="opacity-100 translate-y-0 sm:scale-100"
                            leaveTo="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"
                        >
                            <div className="inline-block align-bottom bg-white rounded-lg px-4 pt-5 pb-4 text-left overflow-hidden shadow-xl transform transition-all sm:my-8 sm:align-middle sm:max-w-lg sm:w-full sm:p-6">
                                <div>

                                    <div className="mx-auto flex items-center justify-center h-12 w-12 rounded-full bg-green-100">
                                        {!this.state.isConnected &&
                                            <ServerIcon className="h-6 w-6 text-green-600" aria-hidden="true" />
                                        }
                                        {this.state.isConnected &&
                                            <CheckIcon className="h-6 w-6 text-green-600" aria-hidden="true" />
                                        }
                                    </div>
                                    <div className="mt-3 text-center sm:mt-5">
                                        {!this.state.isConnected &&
                                            <Dialog.Title as="h3" className="text-lg leading-6 font-medium text-gray-900">
                                                Connect to a Server
                                            </Dialog.Title>
                                        }
                                        {this.state.isConnected &&
                                            <Dialog.Title as="h3" className="text-lg leading-6 font-medium text-gray-900">
                                                A connection can be established
                                            </Dialog.Title>
                                        }
                                        <div className="mt-2">
                                            {!this.state.isConnected &&
                                                <p className="text-sm text-gray-500">
                                                    You must connect to a server to proceed.
                                                </p>
                                            }
                                            {this.state.isConnected &&
                                                <p className="text-sm text-gray-500">
                                                    Hostname: {this.state.sysInfo.hostname} <br />
                                                    Architecture: {this.state.sysInfo.architecture} <br />
                                                    Available Memory: {this.state.sysInfo.memory} <br />
                                                    OS: {this.state.sysInfo.os}
                                                </p>
                                            }
                                        </div>
                                    </div>
                                </div>
                                {!this.state.isConnected &&
                                    <div>

                                        <label htmlFor="email" className="block text-sm font-medium text-gray-700">
                                            Server URL
                                        </label>
                                        <div className="mt-1 relative rounded-md shadow-sm">
                                            <input
                                                ref={this.initFocus}
                                                type="text"
                                                name="email"
                                                id="email"
                                                value={this.state.url}
                                                className="shadow-sm focus:ring-indigo-500 focus:border-indigo-500 block w-full sm:text-sm border-gray-300 rounded-md"
                                                placeholder="http://127.0.0.1:17415"
                                                onChange={this.handleUrlChange}
                                            />
                                            <div className="absolute inset-y-0 right-0 pr-3 flex items-center pointer-events-none">
                                                <ServerIcon className="h-5 w-5 text-green-500" aria-hidden="true" />
                                            </div>
                                        </div>
                                        <p className="mt-2 text-sm text-gray-500" id="email-description">
                                            Your server address will never be uploaded nor analysed by AID.
                                        </p>
                                    </div>
                                }
                                {!this.state.isConnected && this.state.connecting &&
                                    <div className="mt-5 sm:mt-6 sm:grid sm:grid-cols-2 sm:gap-3 sm:grid-flow-row-dense">
                                        <button type="button" className="disabled w-full inline-flex justify-center rounded-md border border-transparent shadow-sm px-4 py-2 bg-red-600 text-base font-medium text-white hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500 sm:col-start-2 sm:text-sm">
                                            <svg className="animate-spin h-5 w-5 mr-3 text-white" viewBox="0 0 24 24">
                                            <circle className="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="4"></circle>
                                            <path className="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                                            </svg>
                                            Processing
                                        </button>
                                        <button
                                            type="button"
                                            className="mt-3 w-full inline-flex justify-center rounded-md border border-gray-300 shadow-sm px-4 py-2 bg-white text-base font-medium text-gray-700 hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 sm:mt-0 sm:col-start-1 sm:text-sm"
                                        >
                                            Cancel
                                        </button>
                                    </div>
                                }
                                {!this.state.isConnected && !this.state.connecting &&
                                    <div className="mt-5 sm:mt-6 sm:grid sm:grid-cols-2 sm:gap-3 sm:grid-flow-row-dense">
                                        <button
                                            type="button"
                                            className="w-full inline-flex justify-center rounded-md border border-transparent shadow-sm px-4 py-2 bg-indigo-600 text-base font-medium text-white hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 sm:col-start-2 sm:text-sm"
                                            onClick={this.handleConnect}
                                        >
                                            Test Connection
                                        </button>
                                        <button
                                            type="button"
                                            className="mt-3 w-full inline-flex justify-center rounded-md border border-gray-300 shadow-sm px-4 py-2 bg-white text-base font-medium text-gray-700 hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 sm:mt-0 sm:col-start-1 sm:text-sm"
                                        >
                                            Cancel
                                        </button>
                                    </div>
                                }
                                {this.state.isConnected &&
                                    <div className="mt-5 sm:mt-6 sm:grid sm:grid-cols-2 sm:gap-3 sm:grid-flow-row-dense">
                                        <button
                                            type="button"
                                            className="w-full inline-flex justify-center rounded-md border border-transparent shadow-sm px-4 py-2 bg-indigo-600 text-base font-medium text-white hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 sm:col-start-2 sm:text-sm"
                                            onClick={this.handleCompleteConnection}
                                        >
                                            Connect to {this.state.sysInfo.hostname}
                                        </button>
                                        <button
                                            type="button"
                                            className="mt-3 w-full inline-flex justify-center rounded-md border border-gray-300 shadow-sm px-4 py-2 bg-white text-base font-medium text-gray-700 hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 sm:mt-0 sm:col-start-1 sm:text-sm"
                                            onClick={this.handleCancel}
                                        >
                                            Cancel
                                        </button>
                                    </div>
                                }

                            </div>
                        </Transition.Child>
                    </div>
                </Dialog>
            </Transition.Root>
        )
    }

}

export default ConnectPage;