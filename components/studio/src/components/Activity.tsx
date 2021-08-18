const activityItems = [
    { project: 'encodingSolver', commit: '2d89f0c8', environment: 'localhost', time: '1h' },
]
export default function Activity() {
    return (
        <div className="bg-gray-50 pr-4 sm:pr-6 lg:pr-8 lg:flex-shrink-0 lg:border-l lg:border-gray-200 xl:pr-0">
            <div className="pl-6 lg:w-80">
                <div className="pt-6 pb-2">
                    <h2 className="text-sm font-semibold">Activity</h2>
                </div>
                <div>
                    <ul className="divide-y divide-gray-200">
                        {activityItems.map((item) => (
                            <li key={item.commit} className="py-4">
                                <div className="flex space-x-3">
                                    <span className="inline-block h-6 w-6 rounded-full overflow-hidden bg-gray-100">
                                        <svg className="h-full w-full text-gray-300" fill="currentColor" viewBox="0 0 24 24">
                                            <path d="M24 20.993V24H0v-2.996A14.977 14.977 0 0112.004 15c4.904 0 9.26 2.354 11.996 5.993zM16.002 8.999a4 4 0 11-8 0 4 4 0 018 0z" />
                                        </svg>
                                    </span>
                                    <div className="flex-1 space-y-1">
                                        <div className="flex items-center justify-between">
                                            <h3 className="text-sm font-medium">You</h3>
                                            <p className="text-sm text-gray-500">{item.time}</p>
                                        </div>
                                        <p className="text-sm text-gray-500">
                                            No Activities Available.
                                            (It is not being collected for now)
                                        </p>
                                    </div>
                                </div>
                            </li>
                        ))}
                    </ul>
                    <div className="py-4 text-sm border-t border-gray-200">
                        <a href="#" className="text-indigo-600 font-semibold hover:text-indigo-900">
                            View all activity <span aria-hidden="true">&rarr;</span>
                        </a>
                    </div>
                </div>
            </div>
        </div>

    )
}