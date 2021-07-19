// Copyright (c) 2021 Xiaozhe Yao et al.
// 
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT
import Moment from "react-moment"
import { Link } from "react-router-dom"
export default function Analysis(props: any) {
    console.log(props.src)
    return (
        <form className="space-y-8 divide-y divide-gray-200">
            <div className="space-y-8 divide-y divide-gray-200">
                <div>
                    <div>
                        <h3 className="text-lg leading-6 font-medium text-gray-900">Analysis</h3>
                        <p className="mt-1 text-sm text-gray-500">
                            Here's a simple analysis of the request, more detailed analysis are on the way.
                        </p>
                    </div>
                    Your request start at <Moment unix>{props.src.startTime.toString().slice(0, -3)}</Moment>, 
                    and the response returned at <Moment unix>{props.src.endTime.toString().slice(0, -3)}</Moment>. 
                    The overall processing time is {props.src.endTime-props.src.startTime} milliseconds.
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
                    </div>
                </div>
            </div></form>
    )
}