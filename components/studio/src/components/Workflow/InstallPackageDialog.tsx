import { Fragment, useState } from 'react'
import { Dialog, Transition, Disclosure } from '@headlessui/react'
import { ChevronUpIcon, ArrowDownIcon } from '@heroicons/react/outline'
import { restclient } from '../../services/apis'
import Toggles from "../Common/Toggles"

export default function InstallPackagesDialog(props: any) {
    const [buildImage, setBuildImage] = useState(true);
    const [createContainer, setCreateContainer] = useState(true);
    const [readLog, setReadLog] = useState(false);
    const [remoteURL, setRemoteURL] = useState("");
    const [logs, setLogs] = useState("");

    function handleURLChange(e: any) {
        setRemoteURL(e.target.value)
    }

    function streamLog(logid: string) {
        restclient.ws_log(logid, (e: any) => {
            let message = ""
            let logs = e.split(/\r?\n/)
            logs.map(function (each: any) {
                console.log(each)
                if (each) {
                    try {
                        let jsonStr: any = JSON.parse(('{"' + each.replace(/^\s+|\s+$/g, '').replace(/=(?=\s|$)/g, '="" ').replace(/\s+(?=([^"]*"[^"]*")*[^"]*$)/g, '", "').replace(/=/g, '": "') + '"}').replace(/""/g, '"'));
                        console.log(jsonStr)
                        let info = jsonStr.time + " [" + jsonStr.level + "]" + " " + jsonStr.msg
                        message += info
                    } catch (error) {
                        message += each
                    }

                }
            })
            setLogs(message)
        });
    }

    function makeInstall() {
        setLogs(">>> Installing...\n Loading logs...");
        restclient.post('/api/install', {
            'remoteURL': remoteURL
        }).then(function (res: any) {
            if (res.data.message == 'success') {
                let vendor = remoteURL.split("/")[3]
                let name = remoteURL.split("/")[4]

                restclient.get('/api/package/' + vendor + "/" + name).then(function (res: any) {
                    let solvers = res.data.Solvers
                    if (solvers.length == 1) {
                        let solverName = solvers[0]
                        if (buildImage) {
                            restclient.post('/api/mutations', {
                                "operation": "build",
                                "vendorName": vendor,
                                "packageName": name,
                                "solverName": solverName.name
                            }).then(function (res) {
                                console.log(res)
                                streamLog(res.data.logID);
                                setReadLog(true);
                            }).catch(function(err) {
                                console.error(err)
                            })
                        }
                    } else {
                        alert("The package contains several solvers. Hence automatic build is disabled.")
                    }
                })
            }
        })
        setReadLog(true);
    }

    return (
        <Transition.Root show={props.open} as={Fragment}>
            <Dialog as="div" auto-reopen="true" className="fixed z-10 inset-0 overflow-y-auto" onClose={props.onClose}>
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
                        <div className="inline-block align-bottom bg-white rounded-lg px-4 pt-5 pb-4 text-left overflow-hidden shadow-xl transform transition-all lg:my-8 lg:align-middle lg:max-w-lg lg:w-full lg:p-6">
                            <div>
                                <div className="mx-auto flex items-center justify-center h-12 w-12 rounded-full bg-indigo-100">
                                    <ArrowDownIcon className="h-6 w-6 text-indigo-600" aria-hidden="true" />
                                </div>
                                <div className="mt-3 text-center sm:mt-5">
                                    <Dialog.Title as="h3" className="text-lg leading-6 font-medium text-gray-900">
                                        Install Package
                                    </Dialog.Title>
                                    <div className="mt-2">
                                        <label htmlFor="package" className="block text-sm font-medium text-gray-700">
                                            URL or Identifer
                                        </label>
                                        <div className="mt-1 w-full">
                                            <input
                                                type="text"
                                                name="package"
                                                id="package"
                                                className="shadow-sm focus:ring-indigo-500 focus:border-indigo-500 block w-full sm:text-md border-gray-300 rounded-md"
                                                aria-describedby="package-description"
                                                value={remoteURL}
                                                onChange={handleURLChange}
                                            />
                                        </div>
                                        <ul>
                                            <li>
                                                <p className="mt-2 text-sm text-gray-500" id="package-description">
                                                    e.g. https://github.com/aidmodels/detectron
                                                </p>
                                            </li>
                                            <li>
                                                <p className="mt-2 text-sm text-gray-500" id="package-description">
                                                    Note: Expert users are suggested to use the CLI.
                                                </p></li>
                                        </ul>
                                    </div>
                                </div>
                            </div>
                            <div className="mt-5">
                                <div className="w-full max-w-md mx-auto bg-white rounded-2xl">
                                    <Disclosure>
                                        {({ open }) => (
                                            <>
                                                <Disclosure.Button className="flex flex-row-reverse justify-between w-full py-2 text-sm font-medium text-right rounded-lg">
                                                    <ChevronUpIcon
                                                        className={`${open ? '' : 'transform rotate-180'
                                                            } w-5 h-5 text-indigo-500`}
                                                    />
                                                </Disclosure.Button>
                                                <Disclosure.Panel className="pt-4 pb-2 text-sm text-gray-500 inline-block w-full font-medium text-left">
                                                    <div className="flex justify-between">
                                                        Build image after installation
                                                        <Toggles enabled={buildImage} handleEnabled={setBuildImage} />
                                                    </div>
                                                    <div className="flex justify-between mt-2">
                                                        Create container after build
                                                        <Toggles enabled={createContainer} handleEnabled={setCreateContainer} />
                                                    </div>
                                                </Disclosure.Panel>
                                            </>
                                        )}
                                    </Disclosure>
                                </div>
                            </div>
                            <div className="mt-5 sm:mt-4 sm:flex sm:flex-row-reverse">
                                <button
                                    type="button"
                                    className="w-full inline-flex justify-center rounded-md border border-transparent shadow-sm px-4 py-2 bg-indigo-600 text-base font-medium text-white hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 sm:ml-3 sm:w-auto sm:text-sm"
                                    onClick={makeInstall}
                                >
                                    Confirm
                                </button>
                                <button
                                    type="button"
                                    className="mt-3 w-full inline-flex justify-center rounded-md border border-gray-300 shadow-sm px-4 py-2 bg-white text-base font-medium text-gray-700 hover:text-gray-500 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 sm:mt-0 sm:w-auto sm:text-sm"
                                    onClick={props.onClose}
                                >
                                    Cancel
                                </button>
                            </div>
                            {readLog &&
                                <div className="mt-5">
                                    <textarea
                                        rows={8}
                                        className="w-full px-3 py-2 text-gray-700 border-gray-300 rounded-md focus:outline-none resize-y text-xs" value={logs} disabled />
                                </div>
                            }
                        </div>
                    </Transition.Child>
                </div>
            </Dialog>
        </Transition.Root>
    )
}
