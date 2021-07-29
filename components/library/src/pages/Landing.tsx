import MainLayout from '../Layouts/MainLayout'
import { Link } from 'react-router-dom'
import { useState } from 'react'
import FeatureGrid from '../components/Landing/FeatureGrid'
import Screenshots from '../components/Landing/Screenshots'
import LogoCloud from '../components/Landing/LogoCloud'
export default function Landing() {

    const [keyword, setKeyword] = useState("");
    function handleKWChange(e: any) {
        setKeyword(e.target.value)
    }
    return (
        <MainLayout>
                    <div className="pt-10 bg-gray-900 sm:pt-16 lg:pt-8 lg:pb-14 lg:overflow-hidden">
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
                                                            value={keyword}
                                                            onChange={handleKWChange}
                                                            placeholder="Enter the model keywords, e.g. Detection."
                                                            className="block w-full px-4 py-3 rounded-md border-0 text-base text-gray-900 placeholder-gray-500 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-cyan-400 focus:ring-offset-gray-900"
                                                        />
                                                    </div>
                                                    <div className="mt-3 sm:mt-0 sm:ml-3">
                                                        <Link to={`/search?kw=${keyword}`}>
                                                            <button
                                                                type="submit"
                                                                className="block w-full py-3 px-4 rounded-md shadow bg-gradient-to-r from-teal-500 to-cyan-600 text-white font-medium hover:from-teal-600 hover:to-cyan-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-cyan-400 focus:ring-offset-gray-900"
                                                            >
                                                                Search
                                                            </button>
                                                        </Link>
                                                    </div>
                                                </div>
                                                <p className="mt-3 text-sm text-gray-300 sm:mt-4">
                                                    or{' '}
                                                    <a href="/search?kw=" className="font-medium text-white">
                                                        View all models
                                                    </a>
                                                    .
                                                </p>
                                            </form>
                                        </div>
                                    </div>
                                </div>
                                <div className="mt-12 -mb-16 sm:-mb-48 lg:m-0 lg:relative">
                                    <div className="mt-12 -mb-10 sm:-mb-24 lg:-mb-80">
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
                    <Screenshots/>
                    <FeatureGrid/>
                    <LogoCloud/>
        </MainLayout>
    )
}
