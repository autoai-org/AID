import ModelCard from './ModelCard'
import {
    ArrowNarrowLeftIcon,
    ArrowNarrowRightIcon,
} from '@heroicons/react/solid'

export default function ModelList(props: any) {
    return (
        <div>
            <ul>
                {props.models.map((model: any, index: number) => (
                    <li key={'model-list-' + index}>
                        <ModelCard key={'model-card-' + index} model={model}></ModelCard>
                    </li>))}
            </ul>
            {false &&
                <nav className="border-t border-gray-200 px-4 flex items-center justify-between sm:px-0">
                    <div className="-mt-px w-0 flex-1 flex">
                        <a
                            href="/"
                            className="border-t-2 border-transparent pt-4 pr-1 inline-flex items-center text-sm font-medium text-gray-500 hover:text-gray-700 hover:border-gray-300"
                        >
                            <ArrowNarrowLeftIcon className="mr-3 h-5 w-5 text-gray-400" aria-hidden="true" />
                            Previous
                        </a>
                    </div>
                    <div className="hidden md:-mt-px md:flex">
                        <a
                            href="/"
                            className="border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300 border-t-2 pt-4 px-4 inline-flex items-center text-sm font-medium"
                        >
                            1
                        </a>
                        {/* Current: "border-indigo-500 text-indigo-600", Default: "border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300" */}
                        <a
                            href="/"
                            className="border-indigo-500 text-indigo-600 border-t-2 pt-4 px-4 inline-flex items-center text-sm font-medium"
                            aria-current="page"
                        >
                            2
                        </a>

                        <span className="border-transparent text-gray-500 border-t-2 pt-4 px-4 inline-flex items-center text-sm font-medium">
                            ...
                        </span>
                        <a
                            href="/"
                            className="border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300 border-t-2 pt-4 px-4 inline-flex items-center text-sm font-medium"
                        >
                            8
                        </a>
                    </div>
                    <div className="-mt-px w-0 flex-1 flex justify-end">
                        <a
                            href="/"
                            className="border-t-2 border-transparent pt-4 pl-1 inline-flex items-center text-sm font-medium text-gray-500 hover:text-gray-700 hover:border-gray-300"
                        >
                            Next
                            <ArrowNarrowRightIcon className="ml-3 h-5 w-5 text-gray-400" aria-hidden="true" />
                        </a>
                    </div>
                </nav>
            }

        </div>
    )
}