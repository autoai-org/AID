// Copyright (c) 2021 Xiaozhe Yao et al.
// 
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT
import ReactJson from 'react-json-view'
import { Link } from 'react-router-dom'
export default function InferenceResult(props: any) {
    return (
        <form className="space-y-8 divide-y divide-gray-200">
            <div className="space-y-8 divide-y divide-gray-200">
                <div>
                    <div>
                        <h3 className="text-lg leading-6 font-medium text-gray-900">Result</h3>
                        <p className="mt-1 text-sm text-gray-500">
                            The solver returned the following result.
                        </p>
                    </div>
                    <ReactJson src={props.src} />
                </div>
                <div className="pt-5">
                    <div className="flex justify-end">
                    <Link to="/">
                        <button
                            type="button"
                            className="bg-white py-2 px-4 border border-gray-300 rounded-md shadow-sm text-sm font-medium text-gray-700 hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
                        >
                            Close
                        </button>
                        </Link>
                        <button
                            onClick={props.onNextStep}
                            type="button"
                            className="ml-3 inline-flex justify-center py-2 px-4 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
                        >
                            Next
                        </button>
                    </div>
                </div>
            </div></form>
    )
}