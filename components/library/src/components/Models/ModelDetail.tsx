import MainLayout from "../../Layouts/MainLayout"
import {
    PaperClipIcon,
    UserIcon,
} from '@heroicons/react/solid'
import { useParams } from "react-router-dom"
import { getModelInfo } from '../../services/api'
import { useEffect, useState } from "react"
import ReactMarkdown from 'react-markdown'
import Moment from 'react-moment';
import { getInitials } from '../../services/utility'
function classNames(...classes: any) {
    return classes.filter(Boolean).join(' ')
}
export default function ModelDetail(props: any) {
    let { vendor } = useParams<{ vendor: string }>();
    let { name } = useParams<{ name: string }>();

    const [detailedInfo, setDetailedInfo] = useState<any>({})
    const [loaded, setLoaded] = useState(false)
    
    useEffect(() => {
        getModelInfo(vendor, name).then(function (res: any) {
            setDetailedInfo(res.data)
            setLoaded(true)
        }).catch(function (err: any) {
            console.log(err)
        })
    }, [vendor, name])
    return (

        <MainLayout>
            <div className="max-w-3xl mx-auto px-4 sm:px-6 md:flex md:items-center md:justify-between md:space-x-5 lg:max-w-7xl lg:px-8 pt-4">
                <div className="flex items-center space-x-5">
                    <div className="flex-shrink-0">
                        <div className="relative">
                            <span className="inline-flex items-center justify-center h-14 w-14 rounded-full bg-gray-500">
                                <span className="text-xl font-medium leading-none text-white">{getInitials(name)}</span>
                            </span>

                            <span className="absolute inset-0 shadow-inner rounded-full" aria-hidden="true" />
                        </div>
                    </div>
                    <div>
                        <h1 className="text-2xl font-bold text-gray-900">{name}</h1>
                        <p className="text-sm font-medium text-gray-500">
                            Uploaded by {' '}
                            <a href="/" className="text-gray-900">
                                {vendor}
                            </a>
                        </p>
                    </div>
                </div>

                <div className="mt-6 flex flex-col-reverse justify-stretch space-y-4 space-y-reverse sm:flex-row-reverse sm:justify-end sm:space-x-reverse sm:space-y-0 sm:space-x-3 md:mt-0 md:flex-row md:space-x-3">
                    <a
                        href="/"
                        type="button"
                        className="inline-flex items-center justify-center px-4 py-2 border border-gray-300 shadow-sm text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-offset-gray-100 focus:ring-blue-500"
                    >
                        Close
                    </a>
                    <button
                        type="button"
                        className="inline-flex items-center justify-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-offset-gray-100 focus:ring-blue-500"
                    >
                        Install Locally
                    </button>
                </div>
            </div>
            {loaded &&
                <div className="mt-8 max-w-3xl mx-auto grid grid-cols-1 gap-6 sm:px-6 lg:max-w-7xl lg:grid-flow-col-dense lg:grid-cols-3">
                    <div className="space-y-6 lg:col-start-1 lg:col-span-2">
                        {/* Description list*/}
                        <section aria-labelledby="applicant-information-title">
                            <div className="bg-white shadow sm:rounded-lg">
                                <div className="px-4 py-5 sm:px-6">
                                    <h2 id="applicant-information-title" className="text-lg leading-6 font-medium text-gray-900">
                                        Package Information
                                    </h2>
                                    <p className="mt-1 max-w-2xl text-sm text-gray-500">Details about the package.</p>
                                </div>
                                <div className="border-t border-gray-200 px-4 py-5 sm:px-6">
                                    <dl className="grid grid-cols-1 gap-x-4 gap-y-8 sm:grid-cols-2">
                                        <div className="sm:col-span-1">
                                            <dt className="text-sm font-medium text-gray-500">Tagline</dt>
                                            <dd className="mt-1 text-sm text-gray-900">{detailedInfo.configuration.package.tagline}</dd>
                                        </div>
                                        <div className="sm:col-span-1">
                                            <dt className="text-sm font-medium text-gray-500">GPU Suggested?</dt>
                                            <dd className="mt-1 text-sm text-gray-900">{detailedInfo.configuration.package.gpu ? 'Yes' : 'No'}</dd>
                                        </div>
                                        <div className="sm:col-span-1">
                                            <dt className="text-sm font-medium text-gray-500">CI Status</dt>
                                            <dd className="mt-1 text-sm text-gray-900">
                                                <img alt="CI_STATUS_UNKNOWN" src={"https://github.com/" + vendor + "/" + name + "/actions/workflows/aid-ci.yml/badge.svg"}></img>
                                            </dd>
                                        </div>
                                        <div className="sm:col-span-2 prose">
                                            <dt className="text-sm font-medium text-gray-500">README</dt>
                                            <ReactMarkdown>{detailedInfo.readme}</ReactMarkdown>
                                        </div>
                                        <div className="sm:col-span-2">
                                            <dt className="text-sm font-medium text-gray-500">Pretrained Weights</dt>
                                            <dd className="mt-1 text-sm text-gray-900">
                                                <ul className="border border-gray-200 rounded-md divide-y divide-gray-200">
                                                    {detailedInfo.pretrained.models.map((model: any) => (
                                                        <li
                                                            key={model.name}
                                                            className="pl-3 pr-4 py-3 flex items-center justify-between text-sm"
                                                        >
                                                            <div className="w-0 flex-1 flex items-center">
                                                                <PaperClipIcon className="flex-shrink-0 h-5 w-5 text-gray-400" aria-hidden="true" />
                                                                <span className="ml-2 flex-1 w-0 truncate">{model.name}</span>
                                                            </div>
                                                            <div className="ml-4 flex-shrink-0">
                                                                <a href={model.url} className="font-medium text-blue-600 hover:text-blue-500">
                                                                    Download
                                                                </a>
                                                            </div>
                                                        </li>
                                                    ))}
                                                </ul>
                                            </dd>
                                        </div>
                                    </dl>
                                </div>
                            </div>
                        </section>
                    </div>
                    <section aria-labelledby="timeline-title" className="lg:col-start-3 lg:col-span-1">
                        <div className="bg-white px-4 py-5 shadow sm:rounded-lg sm:px-6">
                            <h2 id="timeline-title" className="text-lg font-medium text-gray-900">
                                Commits History
                            </h2>

                            {/* Activity Feed */}
                            <div className="mt-6 flow-root">
                                <ul className="-mb-8">
                                    {detailedInfo.commits.map((item: any, itemIdx: any) => (
                                        <li key={itemIdx}>
                                            <div className="relative pb-8">
                                                {itemIdx !== detailedInfo.commits.length - 1 ? (
                                                    <span className="absolute top-4 left-4 -ml-px h-full w-0.5 bg-gray-200" aria-hidden="true" />
                                                ) : null}
                                                <div className="relative flex space-x-3">
                                                    <div>
                                                        <span
                                                            className={classNames(
                                                                'bg-gray-400 h-8 w-8 rounded-full flex items-center justify-center ring-8 ring-white'
                                                            )}
                                                        >
                                                            <UserIcon className="w-5 h-5 text-white" aria-hidden="true" />
                                                        </span>
                                                    </div>
                                                    <div className="min-w-0 flex-1 pt-1.5 flex justify-between space-x-4">
                                                        <div>
                                                            <p className="text-sm text-gray-500">
                                                                {item.commit.author.name}{' '}
                                                            </p>
                                                            <p className="font-sm text-gray-900">
                                                                {item.commit.message}
                                                            </p>

                                                        </div>
                                                        <div className="text-right text-sm whitespace-nowrap text-gray-500">
                                                            <time dateTime={item.commit.author.date}><Moment fromNow>{item.commit.author.date}</Moment></time>
                                                        </div>
                                                    </div>
                                                </div>
                                            </div>
                                        </li>
                                    ))}
                                </ul>
                            </div>
                            <div className="mt-6 flex flex-col justify-stretch">
                                <button
                                    type="button"
                                    className="inline-flex items-center justify-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
                                >
                                    View on GitHub
                                </button>
                            </div>
                        </div>
                    </section>
                </div>
            }
        </MainLayout>
    )
}