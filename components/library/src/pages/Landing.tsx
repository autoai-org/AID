import MainLayout from '../Layouts/MainLayout'
export default function Example() {
    return (
        <MainLayout>
            <main>
                <div className="pt-10 bg-gray-900 sm:pt-16 lg:pt-20 lg:pb-20 lg:overflow-hidden">
                    <div className="mx-auto max-w-7xl lg:px-8">
                        <div className="lg:grid lg:grid-cols-2 lg:gap-8">
                            <div className="mx-auto max-w-md px-4 sm:max-w-2xl sm:px-6 sm:text-center lg:px-0 lg:text-left lg:flex lg:items-center">
                                <div className="lg:py-24">
                                    <h1 className="mt-4 text-4xl tracking-tight font-extrabold text-white sm:mt-5 sm:text-6xl lg:mt-6 xl:text-6xl">
                                        <span className="block">A better way to</span>
                                        <span className="block text-indigo-400">
                                            access machine learning
                                        </span>
                                    </h1>
                                    <p className="text-base text-gray-300 sm:text-xl lg:text-lg xl:text-xl">
                                        Search, download and run machine learning algorithms within a few clicks/commands.
                                    </p>
                                    <div className="mt-10 sm:mt-12">
                                        <form action="#" className="sm:max-w-xl sm:mx-auto lg:mx-0">
                                            <div className="sm:flex">
                                                <div className="min-w-0 flex-1">
                                                    <label htmlFor="keywords" className="sr-only">
                                                        Model Keywords
                                                    </label>
                                                    <input
                                                        id="keywords"
                                                        type="text"
                                                        placeholder="Enter the model keywords, e.g. Detection."
                                                        className="block w-full px-4 py-3 rounded-md border-0 text-base text-gray-900 placeholder-gray-500 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-cyan-400 focus:ring-offset-gray-900"
                                                    />
                                                </div>
                                                <div className="mt-3 sm:mt-0 sm:ml-3">
                                                    <button
                                                        type="submit"
                                                        className="block w-full py-3 px-4 rounded-md shadow bg-gradient-to-r from-teal-500 to-cyan-600 text-white font-medium hover:from-teal-600 hover:to-cyan-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-cyan-400 focus:ring-offset-gray-900"
                                                    >
                                                        Search
                                                    </button>
                                                </div>
                                            </div>
                                            <p className="mt-3 text-sm text-gray-300 sm:mt-4">
                                                or{' '}
                                                <a href="#" className="font-medium text-white">
                                                    View all models
                                                </a>
                                                .
                                            </p>
                                        </form>
                                    </div>
                                </div>
                            </div>
                            <div className="mt-12 -mb-16 sm:-mb-48 lg:m-0 lg:relative">
                                <div className="mx-auto max-w-md px-4 sm:max-w-2xl sm:px-6 lg:max-w-none lg:px-0">
                                    {/* Illustration taken from Lucid Illustrations: https://lucid.pixsellz.io/ */}
                                    <img
                                        className="w-full lg:absolute lg:inset-y-0 lg:left-0 lg:h-full lg:w-auto lg:max-w-none"
                                        src="https://tailwindui.com/img/component-images/cloud-illustration-teal-cyan.svg"
                                        alt=""
                                    />
                                </div>
                            </div>
                        </div>
                    </div>
                </div>

                {/* Feature section with screenshot */}
                <div className="relative bg-gray-50 pt-16 sm:pt-24 lg:pt-32">
                    <div className="mx-auto max-w-md px-4 text-center sm:px-6 sm:max-w-3xl lg:px-8 lg:max-w-7xl">
                        <div>
                            <h2 className="text-base font-semibold tracking-wider text-cyan-600 uppercase">Open Source</h2>
                            <p className="mt-2 text-3xl font-extrabold text-gray-900 tracking-tight sm:text-4xl">
                                AID is open source and available on{' '} GitHub
                            </p>
                            <p className="mt-5 max-w-prose mx-auto text-xl text-gray-500">
                                If you have sensitive data or models, it is best to save them on your own server.
                            </p>
                        </div>
                        <div className="mt-12 -mb-10 sm:-mb-24 lg:-mb-80">
                            <img
                                className="rounded-lg shadow-xl ring-1 ring-black ring-opacity-5"
                                src="/assets/screenshots.png"
                                alt=""
                            />
                        </div>
                    </div>
                </div>
            </main>
        </MainLayout>
    )
}
