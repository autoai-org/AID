import { Menu } from '@headlessui/react'
import { ALL_SOLVERS } from '../services/apis/queries'
import { restclient } from '../services/apis';
import { setIsLoading } from '../services/store/connectivity/server';
import Toggles from "../components/Common/Toggles";
import { useDispatch } from 'react-redux'
import { useHistory } from 'react-router-dom';
import {
    ChevronDownIcon,
    ChevronRightIcon,
    SortAscendingIcon,
    StarIcon,
    PlusIcon
} from '@heroicons/react/solid'
import Moment from 'react-moment';
import { useEffect, useState } from 'react';
import { constants } from 'buffer';

function classNames(...classes: any) {
    return classes.filter(Boolean).join(' ')
}

export default function SolverColumn() {
    const dispatch = useDispatch();
    const [solvers, setSolvers] = useState<object[]>([]);
    const history = useHistory();
    useEffect(() => {
        dispatch(setIsLoading(true))
        restclient.query(ALL_SOLVERS).then((res: any) => {
            setSolvers(res.data.data.solvers)
            localStorage.setItem("aid-solvers", JSON.stringify(res.data.data.solvers))
        }).finally(() => {
            dispatch(setIsLoading(false))
        })
    }, [])

    const handleToggling = () =>{

    } 

    return (
        <div className="bg-white lg:min-w-0 lg:flex-1">
            <div className="pl-4 pr-6 pt-4 pb-4 border-b border-t border-gray-200 sm:pl-6 lg:pl-8 xl:pl-6 xl:pt-6 xl:border-t-0">
                <div className="flex items-center">
                    <h1 className="flex-1 text-lg font-medium">Solvers & Services</h1>
                    <Menu as="div" className="relative">
                        <Menu.Button className="w-full bg-white border border-gray-300 rounded-md shadow-sm px-4 py-2 inline-flex justify-center text-sm font-medium text-gray-700 hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">
                            <SortAscendingIcon className="mr-3 h-5 w-5 text-gray-400" aria-hidden="true" />
                            Sort
                            <ChevronDownIcon className="ml-2.5 -mr-1.5 h-5 w-5 text-gray-400" aria-hidden="true" />
                        </Menu.Button>
                        <Menu.Items className="origin-top-right z-10 absolute right-0 mt-2 w-56 rounded-md shadow-lg bg-white ring-1 ring-black ring-opacity-5 focus:outline-none">
                            <div className="py-1">
                                <Menu.Item>
                                    {({ active }) => (
                                        <a
                                            href="#"
                                            className={classNames(
                                                active ? 'bg-gray-100 text-gray-900' : 'text-gray-700',
                                                'block px-4 py-2 text-sm'
                                            )}
                                        >
                                            Name
                                        </a>
                                    )}
                                </Menu.Item>
                                <Menu.Item>
                                    {({ active }) => (
                                        <a
                                            href="#"
                                            className={classNames(
                                                active ? 'bg-gray-100 text-gray-900' : 'text-gray-700',
                                                'block px-4 py-2 text-sm'
                                            )}
                                        >
                                            Date created
                                        </a>
                                    )}
                                </Menu.Item>
                            </div>
                        </Menu.Items>
                    </Menu>
                </div>
            </div>
            {solvers.length === 0 &&
                <div className="text-center">
                    <svg
                        className="mx-auto h-12 w-12 text-gray-400"
                        fill="none"
                        viewBox="0 0 24 24"
                        stroke="currentColor"
                        aria-hidden="true"
                    >
                        <path
                            vectorEffect="non-scaling-stroke"
                            strokeLinecap="round"
                            strokeLinejoin="round"
                            strokeWidth={2}
                            d="M9 13h6m-3-3v6m-9 1V7a2 2 0 012-2h6l2 2h6a2 2 0 012 2v8a2 2 0 01-2 2H5a2 2 0 01-2-2z"
                        />
                    </svg>
                    <h3 className="mt-2 text-sm font-medium text-gray-900">No solvers found</h3>
                    <p className="mt-1 text-sm text-gray-500">Get started by installing a new package.</p>
                    <div className="mt-6">
                        <button
                            type="button"
                            className="inline-flex items-center px-4 py-2 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
                        >
                            <PlusIcon className="-ml-1 mr-2 h-5 w-5" aria-hidden="true" />
                            New Project
                        </button>
                    </div>
                </div>
            }
            {solvers.length > 0 &&
                <ul className="relative z-0 divide-y divide-gray-200 border-b border-gray-200">
                    {solvers.map((solver: any) => (
    
                        <li
                            key={solver.name}
                            className="relative pl-4 pr-6 py-5 hover:bg-gray-50 sm:py-6 sm:pl-6 lg:pl-8 xl:pl-6"
                        >
            
                            <div className="flex items-center justify-between space-x-4">
                                {/* Repo name and link */}
                                <div className="min-w-0 space-y-3">
                                    <div className="flex items-center space-x-3">
                                        <span
                                            className={classNames(
                                                solver.status ? 'bg-green-100' : 'bg-gray-100',
                                                'h-4 w-4 rounded-full flex items-center justify-center'
                                            )}
                                            aria-hidden="true"
                                        >
                                            <span
                                                className={classNames(
                                                    solver.status ? 'bg-green-400' : 'bg-gray-400',
                                                    'h-2 w-2 rounded-full'
                                                )}
                                            />
                                        </span>
                                        <span className="block">
                                            <h2 className="text-sm font-medium">
                                                <a href={"/details/"+solver.uid}>
                                                    <span className="absolute inset-0" aria-hidden="true" />
                                                    {solver.name}{' '}
                                                    <span className="sr-only">{solver.repository.name ? 'Running' : 'Not running'}</span>
                                                </a>
                                            </h2>
                                        </span>
                                    </div>
                                    <a href={solver.repository.remote_url} target="_blank" className="relative group flex items-center space-x-2.5">
                                        <svg
                                            className="flex-shrink-0 w-5 h-5 text-gray-400 group-hover:text-gray-500"
                                            viewBox="0 0 18 18"
                                            fill="none"
                                            xmlns="http://www.w3.org/2000/svg"
                                            aria-hidden="true"
                                        >
                                            <path
                                                fillRule="evenodd"
                                                clipRule="evenodd"
                                                d="M8.99917 0C4.02996 0 0 4.02545 0 8.99143C0 12.9639 2.57853 16.3336 6.15489 17.5225C6.60518 17.6053 6.76927 17.3277 6.76927 17.0892C6.76927 16.8762 6.76153 16.3104 6.75711 15.5603C4.25372 16.1034 3.72553 14.3548 3.72553 14.3548C3.31612 13.316 2.72605 13.0395 2.72605 13.0395C1.9089 12.482 2.78793 12.4931 2.78793 12.4931C3.69127 12.5565 4.16643 13.4198 4.16643 13.4198C4.96921 14.7936 6.27312 14.3968 6.78584 14.1666C6.86761 13.5859 7.10022 13.1896 7.35713 12.965C5.35873 12.7381 3.25756 11.9665 3.25756 8.52116C3.25756 7.53978 3.6084 6.73667 4.18411 6.10854C4.09129 5.88114 3.78244 4.96654 4.27251 3.72904C4.27251 3.72904 5.02778 3.48728 6.74717 4.65082C7.46487 4.45101 8.23506 4.35165 9.00028 4.34779C9.76494 4.35165 10.5346 4.45101 11.2534 4.65082C12.9717 3.48728 13.7258 3.72904 13.7258 3.72904C14.217 4.96654 13.9082 5.88114 13.8159 6.10854C14.3927 6.73667 14.7408 7.53978 14.7408 8.52116C14.7408 11.9753 12.6363 12.7354 10.6318 12.9578C10.9545 13.2355 11.2423 13.7841 11.2423 14.6231C11.2423 15.8247 11.2313 16.7945 11.2313 17.0892C11.2313 17.3299 11.3937 17.6097 11.8501 17.522C15.4237 16.3303 18 12.9628 18 8.99143C18 4.02545 13.97 0 8.99917 0Z"
                                                fill="currentcolor"
                                            />
                                        </svg>
                                        <span className="text-sm text-gray-500 group-hover:text-gray-900 font-medium truncate">
                                            {solver.repository.name}
                                        </span>
                                    </a>
                                </div>
                                <div className="sm:hidden">
                                    <ChevronRightIcon className="h-5 w-5 text-gray-400" aria-hidden="true" />
                                </div>
                                {/* Repo meta info */}
                                <div className="hidden sm:flex flex-col flex-shrink-0 items-end space-y-3">
                                    <p className="flex items-center space-x-4">

                                        <Toggles enabled={solver.status} handleEnabled={handleToggling}/>

                                        <button
                                            className="relative bg-white rounded-full focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
                                            type="button"
                                        >
                                            <span className="sr-only">
                                                {solver.repository.vendor ? 'Add to favorites' : 'Remove from favorites'}
                                            </span>
                                            <StarIcon
                                                className={classNames(
                                                    solver.repository.name
                                                        ? 'text-yellow-300 hover:text-yellow-400'
                                                        : 'text-gray-300 hover:text-gray-400',
                                                    'h-5 w-5'
                                                )}
                                                aria-hidden="true"
                                            />
                                        </button>
                                    </p>
                                    <p className="flex text-gray-500 text-sm space-x-2">
                                        <span>By {solver.repository.vendor}</span>
                                        <span aria-hidden="true">&middot;</span>
                                        <span>Installed <Moment fromNow>{solver.repository.created_at}</Moment></span>
                                    </p>
                                </div>
                            </div>
                         
                        </li>
                        
                    ))}
                </ul>
            }
        </div>
    )
}