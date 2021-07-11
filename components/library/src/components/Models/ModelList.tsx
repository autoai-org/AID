// Copyright (c) 2021 Xiaozhe Yao et al.
// 
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

import {
    CheckCircleIcon,
    ChevronRightIcon,
    MailIcon,
} from '@heroicons/react/solid'


export default function ModelList(props: any) {
    <ul className="mt-5 border-t border-gray-200 divide-y divide-gray-200 sm:mt-0 sm:border-t-0" role="list">
        {props.models.map((model: any) => (
            <li key={model.email}>
                <a href="#" className="group block">
                    <div className="flex items-center py-5 px-4 sm:py-6 sm:px-0">
                        <div className="min-w-0 flex-1 flex items-center">
                            <div className="flex-shrink-0">
                                <img
                                    className="h-12 w-12 rounded-full group-hover:opacity-75"
                                    src={model.imageUrl}
                                    alt=""
                                />
                            </div>
                            <div className="min-w-0 flex-1 px-4 md:grid md:grid-cols-2 md:gap-4">
                                <div>
                                    <p className="text-sm font-medium text-purple-600 truncate">{model.name}</p>
                                    <p className="mt-2 flex items-center text-sm text-gray-500">
                                        <MailIcon className="flex-shrink-0 mr-1.5 h-5 w-5 text-gray-400" aria-hidden="true" />
                                        <span className="truncate">{model.email}</span>
                                    </p>
                                </div>
                                <div className="hidden md:block">
                                    <div>
                                        <p className="text-sm text-gray-900">
                                            Applied on <time dateTime={model.appliedDatetime}>{model.applied}</time>
                                        </p>
                                        <p className="mt-2 flex items-center text-sm text-gray-500">
                                            <CheckCircleIcon
                                                className="flex-shrink-0 mr-1.5 h-5 w-5 text-green-400"
                                                aria-hidden="true"
                                            />
                                            {model.status}
                                        </p>
                                    </div>
                                </div>
                            </div>
                        </div>
                        <div>
                            <ChevronRightIcon
                                className="h-5 w-5 text-gray-400 group-hover:text-gray-700"
                                aria-hidden="true"
                            />
                        </div>
                    </div>
                </a>
            </li>
        ))}
    </ul>
}