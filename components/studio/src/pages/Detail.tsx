import { useState, useEffect } from 'react';
import ReactMarkdown from 'react-markdown'
import {
  PaperClipIcon,
  UserIcon,
} from '@heroicons/react/solid'
import { Fragment } from 'react'
import { Popover, Transition } from '@headlessui/react'
import BackToHomepage from '../components/BacktoHompage'
import { setIsLoading } from '../services/store/connectivity/server'
import { useDispatch } from 'react-redux'
import { restclient } from '../services/apis';
import Moment from 'react-moment';
import { getInitials } from '../services/utilities/initials'
import { serverEndpoint } from '../services/apis'
import toml from 'toml'

function classNames(...classes: any) {
  return classes.filter(Boolean).join(' ')
}

interface commit {
  author: string,
  message: string,
  when: string
}

interface container {
  uid: string,
  running: boolean,
  created_at: string,
}

interface SolverInfo {
  solvername: string;
  description: string;
  vendorname: string;
  remoteURL: string;
  commits: commit[];
  pretrained: any[];
  containers: container[];
}

export default function Details(props: any) {
  let defaultSolverInfoInfo: SolverInfo = {
    solvername: "An Awesome Solver!",
    description: "Awesome Description",
    vendorname: "Awesome Company",
    remoteURL: "",
    pretrained: [],
    commits: [],
    containers: [],
  }
  const [solverInfo, setSolverInfo] = useState(defaultSolverInfoInfo);
  const [loaded, setLoaded] = useState(false);
  const dispatch = useDispatch();

  useEffect(() => {
    dispatch(setIsLoading(true))
    restclient.get(serverEndpoint+ "/solver/" + props.match.params.solverID).then(function (res: any) {
      let pretrained = toml.parse(res.data.pretrained)
      res.data.pretrained = pretrained.models
      console.log(res.data)
      setSolverInfo(res.data)
      
    }).catch(function (err) {
    }).finally(function () {
      dispatch(setIsLoading(false))
      setLoaded(true)
    })
  }, [])

  function loadContainerCreation() {
    console.log("-===")
  }

  return (
    <div className="min-h-screen bg-gray-100">
      <BackToHomepage />

      <main className="py-10">
        {/* Page header */}
        <div className="max-w-3xl mx-auto px-4 sm:px-6 md:flex md:items-center md:justify-between md:space-x-5 lg:max-w-7xl lg:px-8">
          <div className="flex items-center space-x-5">
            <div className="flex-shrink-0">
              <div className="relative">
                <span className="inline-flex items-center justify-center h-14 w-14 rounded-full bg-gray-500">
                  <span className="text-xl font-medium leading-none text-white">{getInitials(solverInfo.solvername)}</span>
                </span>
                <span className="absolute inset-0 shadow-inner rounded-full" aria-hidden="true" />
              </div>
            </div>
            <div>
              <h1 className="text-2xl font-bold text-gray-900">{solverInfo.solvername}</h1>
              <p className="text-sm font-medium text-gray-500">
                Uploaded by{' '}
                <a href="#" className="text-gray-900">
                  {solverInfo.vendorname}
                </a>
              </p>
            </div>
          </div>
          <div className="mt-6 flex flex-col-reverse justify-stretch space-y-4 space-y-reverse sm:flex-row-reverse sm:justify-end sm:space-x-reverse sm:space-y-0 sm:space-x-3 md:mt-0 md:flex-row md:space-x-3">
            <a target="_blank" href={solverInfo.remoteURL}>
              <button
                type="button"
                className="inline-flex items-center justify-center px-4 py-2 border border-gray-300 shadow-sm text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-offset-gray-100 focus:ring-blue-500"
              >
                Homepage
              </button>
            </a>
            <Popover className="relative">
              {({ open }) => (
                <>
                  <Popover.Button
                    className={classNames(
                      open ? 'text-gray-900' : 'text-gray-500',
                      'group bg-white rounded-md inline-flex items-center text-base font-medium hover:text-gray-900 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500'
                    )}
                  >
                    <div
                      className="inline-flex items-center justify-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-offset-gray-100 focus:ring-blue-500"
                    >
                      Inference
                    </div>
                  </Popover.Button>
                  <Transition
                    show={open}
                    as={Fragment}
                    enter="transition ease-out duration-200"
                    enterFrom="opacity-0 translate-y-1"
                    enterTo="opacity-100 translate-y-0"
                    leave="transition ease-in duration-150"
                    leaveFrom="opacity-100 translate-y-0"
                    leaveTo="opacity-0 translate-y-1"
                  >
                    <Popover.Panel
                      static
                      className="absolute z-10 left-1/2 transform -translate-x-1/2 mt-3 px-2 w-screen max-w-md sm:px-0"
                    >
                      <div className="rounded-lg shadow-lg ring-1 ring-black ring-opacity-5 overflow-hidden">
                        <div className="relative grid gap-6 bg-white px-5 py-6 sm:gap-8 sm:p-8">
                          {solverInfo.containers.map((item) => (
                            <a
                              key={item.uid}
                              href={"/inference/"+item.uid}
                              className="-m-3 p-3 flex items-start rounded-lg hover:bg-gray-50 transition ease-in-out duration-150"
                            >
                              {item.uid}
                            </a>
                          ))}
                        </div>
                        <div className="px-5 py-5 bg-gray-50 space-y-6 sm:flex sm:space-y-0 sm:space-x-10 sm:px-8">
                            Above are the running solvers. 
                            <div className="rounded items-center justify-center border border-transparent bg-blue-600 hover:bg-blue-700 text-white ml-2" onClick={loadContainerCreation}>Create</div>
                        </div>
                      </div>
                    </Popover.Panel>
                  </Transition>
                </>)}
            </Popover>
          </div>
        </div>

        <div className="mt-8 max-w-3xl mx-auto grid grid-cols-1 gap-6 sm:px-6 lg:max-w-7xl lg:grid-flow-col-dense lg:grid-cols-3">
          <div className="space-y-6 lg:col-start-1 lg:col-span-2">
            {/* Description list*/}
            <section aria-labelledby="applicant-information-title">
              <div className="bg-white shadow sm:rounded-lg">
                <div className="px-4 py-5 sm:px-6">
                  <h2 id="applicant-information-title" className="text-lg leading-6 font-medium text-gray-900">
                    Model Description
                  </h2>
                  <p className="mt-1 max-w-2xl text-sm text-gray-500">Model details.</p>
                </div>
                <div className="border-t border-gray-200 px-4 py-5 sm:px-6">
                  <dl className="grid grid-cols-1 gap-x-4 gap-y-8 sm:grid-cols-2">
                    <div className="sm:col-span-1">
                      <dt className="text-sm font-medium text-gray-500">Made for</dt>
                      <dd className="mt-1 text-sm text-gray-900">Image Encoding</dd>
                    </div>
                    <div className="sm:col-span-1">
                      <dt className="text-sm font-medium text-gray-500">GPU Necessary?</dt>
                      <dd className="mt-1 text-sm text-gray-900">NO</dd>
                    </div>
                    <div className="sm:col-span-1">
                      <dt className="text-sm font-medium text-gray-500">Required Input</dt>
                      <dd className="mt-1 text-sm text-gray-900">Image</dd>
                    </div>
                    <div className="sm:col-span-1">
                      <dt className="text-sm font-medium text-gray-500">Frameworks</dt>
                      <dd className="mt-1 text-sm text-gray-900">TensorFlow</dd>
                    </div>
                    <div className="sm:col-span-1">
                      <dt className="text-sm font-medium text-gray-500">CI Status</dt>
                      <dd className="mt-1 text-sm text-gray-900">
                        <img src="https://github.com/aidmodels/image_encoding/actions/workflows/aid-ci.yml/badge.svg"></img>
                      </dd>
                    </div>
                    <div className="sm:col-span-2 prose">
                      <dt className="text-sm font-medium text-gray-500">About</dt>
                      <ReactMarkdown>{solverInfo.description}</ReactMarkdown>
                    </div>
                    <div className="sm:col-span-2">
                      <dt className="text-sm font-medium text-gray-500">Attachments</dt>
                      <dd className="mt-1 text-sm text-gray-900">
                        <ul className="border border-gray-200 rounded-md divide-y divide-gray-200">
                          {solverInfo.pretrained.map((attachment) => (
                            <li
                              key={attachment.name}
                              className="pl-3 pr-4 py-3 flex items-center justify-between text-sm"
                            >
                              <div className="w-0 flex-1 flex items-center">
                                <PaperClipIcon className="flex-shrink-0 h-5 w-5 text-gray-400" aria-hidden="true" />
                                <span className="ml-2 flex-1 w-0 truncate">{attachment.name}</span>
                              </div>
                              <div className="ml-4 flex-shrink-0">
                                <a href={attachment.url} className="font-medium text-blue-600 hover:text-blue-500">
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
                <div>
                  <a
                    href="#"
                    className="block bg-gray-50 text-sm font-medium text-gray-500 text-center px-4 py-4 hover:text-gray-700 sm:rounded-b-lg"
                  >
                    Read in Model Hub
                  </a>
                </div>
              </div>
            </section>
          </div>
          <section aria-labelledby="timeline-title" className="lg:col-start-3 lg:col-span-1">
            <div className="bg-white px-4 py-5 shadow sm:rounded-lg sm:px-6">
              <h2 id="timeline-title" className="text-lg font-medium text-gray-900">
                Timeline
              </h2>

              {/* Activity Feed */}
              <div className="mt-6 flow-root">
                <ul className="-mb-8">
                  {solverInfo.commits.map((item, itemIdx) => (
                    <li key={itemIdx}>
                      <div className="relative pb-8">
                        {itemIdx !== solverInfo.commits.length - 1 ? (
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
                                {item.author}{' '}
                              </p>
                              <p className="font-sm text-gray-900">
                                {item.message}
                              </p>

                            </div>
                            <div className="text-right text-sm whitespace-nowrap text-gray-500">
                              <time dateTime={item.when}><Moment fromNow>{item.when}</Moment></time>
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
                  View Full Timeline
                </button>
              </div>
            </div>
          </section>
        </div>
      </main>
    </div>
  )
}