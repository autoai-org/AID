// Copyright (c) 2021 Xiaozhe Yao et al.
// 
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT
import React from 'react'
import HTTPClient from "../../services/apis/solver"


interface RequestFormState {
    textValue: string;
}

class RequestForm extends React.Component<any, RequestFormState> {
    constructor(props: any) {
        super(props);
        this.state = {
            textValue: ''
        };
    }
    sendRequest = () => {
        let httpc = new HTTPClient("http://127.0.0.1:17415/running/38a64faa/infer")
        let self = this
        httpc.addPayload("text", this.state.textValue)
        httpc.send().then(function (res) {
            self.props.onReceivedResponse(JSON.parse(res))
        })
    }
    handleChange = (event: any) => {
        this.setState({ textValue: event.target.value });
        console.log(this.state.textValue)
    }
    render() {
        return (
            <form className="space-y-8 divide-y divide-gray-200">
                <div className="space-y-8 divide-y divide-gray-200">
                    <div>
                        <div>
                            <h3 className="text-lg leading-6 font-medium text-gray-900">Request Form</h3>
                            <p className="mt-1 text-sm text-gray-500">
                                The form below will be sent to solver.
                            </p>
                        </div>

                        <div className="mt-6 grid grid-cols-1 gap-y-6 gap-x-4 sm:grid-cols-6">
                            <div className="sm:col-span-6">
                                <label htmlFor="about" className="block text-sm font-medium text-gray-700">
                                    Text
                                </label>
                                <div className="mt-1">
                                    <textarea
                                        id="about"
                                        name="about"
                                        rows={3}
                                        className="shadow-sm focus:ring-indigo-500 focus:border-indigo-500 block w-full sm:text-sm border border-gray-300 rounded-md"
                                        value={this.state.textValue}
                                        onChange={this.handleChange}
                                    />
                                </div>
                            </div>

                            <div className="sm:col-span-6">
                                <label htmlFor="cover-photo" className="block text-sm font-medium text-gray-700">
                                    File
                                </label>
                                <div className="mt-1 flex justify-center px-6 pt-5 pb-6 border-2 border-gray-300 border-dashed rounded-md">
                                    <div className="space-y-1 text-center">
                                        <svg
                                            className="mx-auto h-12 w-12 text-gray-400"
                                            stroke="currentColor"
                                            fill="none"
                                            viewBox="0 0 48 48"
                                            aria-hidden="true"
                                        >
                                            <path
                                                d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-8l-4-4m0 0L8 8m4-4v12"
                                                strokeWidth={2}
                                                strokeLinecap="round"
                                                strokeLinejoin="round"
                                            />
                                        </svg>
                                        <div className="flex text-sm text-gray-600">
                                            <label
                                                htmlFor="file-upload"
                                                className="relative cursor-pointer bg-white rounded-md font-medium text-indigo-600 hover:text-indigo-500 focus-within:outline-none focus-within:ring-2 focus-within:ring-offset-2 focus-within:ring-indigo-500"
                                            >
                                                <span>Upload a file</span>
                                                <input id="file-upload" name="file-upload" type="file" className="sr-only" />
                                            </label>
                                            <p className="pl-1">or drag and drop</p>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
                <div className="pt-5">
                    <div className="flex justify-end">
                        <button
                            type="button"
                            className="bg-white py-2 px-4 border border-gray-300 rounded-md shadow-sm text-sm font-medium text-gray-700 hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
                        >
                            Cancel
                        </button>
                        <button
                            onClick={this.sendRequest}
                            type="button"
                            className="ml-3 inline-flex justify-center py-2 px-4 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
                        >
                            Send
                        </button>
                    </div>
                </div>
            </form>
        )
    }
}

export default RequestForm