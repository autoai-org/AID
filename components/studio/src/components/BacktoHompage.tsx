import {
    ArrowNarrowLeftIcon,
    HomeIcon,
  } from '@heroicons/react/solid'

export default function BackToHomepage() {
    return(
        <header className="bg-white shadow">
            <div className="max-w-7xl mx-auto px-4 sm:px-6">
            <div className="border-t border-gray-200 py-3">
                <nav className="flex" aria-label="Breadcrumb">
                <div className="flex sm:hidden">
                    <a
                    href="/"
                    className="group inline-flex space-x-3 text-sm font-medium text-gray-500 hover:text-gray-700"
                    >
                    <ArrowNarrowLeftIcon
                        className="flex-shrink-0 h-5 w-5 text-gray-400 group-hover:text-gray-600"
                        aria-hidden="true"
                    />
                    <span>Back to Homepage</span>
                    </a>
                </div>
                <div className="hidden sm:block">
                    <ol className="flex items-center space-x-4">
                    <li>
                        <div>
                        <a href="/" className="text-gray-400 hover:text-gray-500">
                            <HomeIcon className="flex-shrink-0 h-5 w-5" aria-hidden="true" />
                            <span className="sr-only">Home</span>
                        </a>
                        </div>
                    </li>

                    </ol>
                </div>
                </nav>
            </div>
         </div>
      </header>

    )
}
