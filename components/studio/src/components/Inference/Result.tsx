// Copyright (c) 2021 Xiaozhe Yao et al.
// 
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT
import ReactJson from 'react-json-view'

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
            </div></form>
    )
}