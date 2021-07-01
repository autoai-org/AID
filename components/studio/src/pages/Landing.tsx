import {
  BadgeCheckIcon,
  CollectionIcon,
} from '@heroicons/react/solid'

import Connection from './Connection'
import { useAppSelector, useAppDispatch } from '../services/store/hooks'
import SolversColumn from '../components/Solvers'
import AIDHeader from '../components/Header'

const activityItems = [
  { project: 'Workcation', commit: '2d89f0c8', environment: 'production', time: '1h' },
  // More items...
]


export default function Landing() {
  const server = useAppSelector((state) => state.connectivity.server)
  const dispatch = useAppDispatch()
  if (server!=="") {
    return (
      <>
        {/* Background color split screen for large screens */}
        <div className="fixed top-0 left-0 w-1/2 h-full bg-white" aria-hidden="true" />
        <div className="fixed top-0 right-0 w-1/2 h-full bg-gray-50" aria-hidden="true" />
        <div className="relative min-h-screen flex flex-col">
          {/* Navbar */}
          
          <AIDHeader/>
          {/* 3 column wrapper */}
          <div className="flex-grow w-full max-w-7xl mx-auto xl:px-8 lg:flex">
            {/* Left sidebar & main wrapper */}
            <div className="flex-1 min-w-0 bg-white xl:flex">
              {/* Account profile */}
              <div className="xl:flex-shrink-0 xl:w-64 xl:border-r xl:border-gray-200 bg-white">
                <div className="pl-4 pr-6 py-6 sm:pl-6 lg:pl-8 xl:pl-0">
                  <div className="flex items-center justify-between">
                    <div className="flex-1 space-y-8">
                      <div className="space-y-8 sm:space-y-0 sm:flex sm:justify-between sm:items-center xl:block xl:space-y-8">
                        {/* Profile */}
                        <div className="flex items-center space-x-3">
                          <div className="flex-shrink-0 h-12 w-12">
                            <img
                              className="h-12 w-12 rounded-full"
                              src="https://images.unsplash.com/photo-1517365830460-955ce3ccd263?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=256&h=256&q=80"
                              alt=""
                            />
                          </div>
                          <div className="space-y-1">
                            <div className="text-sm font-medium text-gray-900">Debbie Lewis</div>
                            <a href="#" className="group flex items-center space-x-2.5">
                              <svg
                                className="h-5 w-5 text-gray-400 group-hover:text-gray-500"
                                aria-hidden="true"
                                fill="currentColor"
                                viewBox="0 0 20 20"
                              >
                                <path
                                  fillRule="evenodd"
                                  d="M10 0C4.477 0 0 4.484 0 10.017c0 4.425 2.865 8.18 6.839 9.504.5.092.682-.217.682-.483 0-.237-.008-.868-.013-1.703-2.782.605-3.369-1.343-3.369-1.343-.454-1.158-1.11-1.466-1.11-1.466-.908-.62.069-.608.069-.608 1.003.07 1.531 1.032 1.531 1.032.892 1.53 2.341 1.088 2.91.832.092-.647.35-1.088.636-1.338-2.22-.253-4.555-1.113-4.555-4.951 0-1.093.39-1.988 1.029-2.688-.103-.253-.446-1.272.098-2.65 0 0 .84-.27 2.75 1.026A9.564 9.564 0 0110 4.844c.85.004 1.705.115 2.504.337 1.909-1.296 2.747-1.027 2.747-1.027.546 1.379.203 2.398.1 2.651.64.7 1.028 1.595 1.028 2.688 0 3.848-2.339 4.695-4.566 4.942.359.31.678.921.678 1.856 0 1.338-.012 2.419-.012 2.747 0 .268.18.58.688.482A10.019 10.019 0 0020 10.017C20 4.484 15.522 0 10 0z"
                                  clipRule="evenodd"
                                />
                              </svg>
                              <span className="text-sm text-gray-500 group-hover:text-gray-900 font-medium">
                                debbielewis
                              </span>
                            </a>
                          </div>
                        </div>
                        {/* Action buttons */}
                        <div className="flex flex-col sm:flex-row xl:flex-col">
                          <button
                            type="button"
                            className="inline-flex items-center justify-center px-4 py-2 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 xl:w-full"
                          >
                            New Project
                          </button>
                          <button
                            type="button"
                            className="mt-3 inline-flex items-center justify-center px-4 py-2 border border-gray-300 shadow-sm text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 sm:mt-0 sm:ml-3 xl:ml-0 xl:mt-3 xl:w-full"
                          >
                            Invite Team
                          </button>
                        </div>
                      </div>
                      {/* Meta info */}
                      <div className="flex flex-col space-y-6 sm:flex-row sm:space-y-0 sm:space-x-8 xl:flex-col xl:space-x-0 xl:space-y-6">
                        <div className="flex items-center space-x-2">
                          <BadgeCheckIcon className="h-5 w-5 text-gray-400" aria-hidden="true" />
                          <span className="text-sm text-gray-500 font-medium">Pro Member</span>
                        </div>
                        <div className="flex items-center space-x-2">
                          <CollectionIcon className="h-5 w-5 text-gray-400" aria-hidden="true" />
                          <span className="text-sm text-gray-500 font-medium">8 Projects</span>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
  
              {/* Projects List */}
              <SolversColumn></SolversColumn>
            </div>
            {/* Activity feed */}
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
                          <img
                            className="h-6 w-6 rounded-full"
                            src="https://images.unsplash.com/photo-1517365830460-955ce3ccd263?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=256&h=256&q=80"
                            alt=""
                          />
                          <div className="flex-1 space-y-1">
                            <div className="flex items-center justify-between">
                              <h3 className="text-sm font-medium">You</h3>
                              <p className="text-sm text-gray-500">{item.time}</p>
                            </div>
                            <p className="text-sm text-gray-500">
                              Deployed {item.project} ({item.commit} in master) to {item.environment}
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
          </div>
        </div>
      </>
    )
  } else {
    return (
      <Connection handler={dispatch}></Connection>
    )
  }
  
}