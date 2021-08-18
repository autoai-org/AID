import { Fragment, useState } from 'react'
import { Dialog, Transition, Disclosure } from '@headlessui/react'
import { InformationCircleIcon, XIcon, ChevronUpIcon } from '@heroicons/react/outline'
import Toggles from "../Common/Toggles"

export default function InstallPackagesDialog(props: any) {
    //const close = props.onClose

    return (
        <Transition.Root show={props.open} as={Fragment}>
            <Dialog as="div" auto-reopen="true" className="fixed z-10 inset-0 overflow-y-auto" onClose={props.onClose}>
                <div className="flex items-end justify-center min-h-screen pt-4 px-4 pb-20 text-center sm:block sm:p-0">
                    <Transition.Child
                        as={Fragment}
                        enter="ease-out duration-300"
                        enterFrom="opacity-0"
                        enterTo="opacity-100"
                        leave="ease-in duration-200"
                        leaveFrom="opacity-100"
                        leaveTo="opacity-0"
                    >
                        <Dialog.Overlay className="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity" />
                    </Transition.Child>

                    {/* This element is to trick the browser into centering the modal contents. */}
                    <span className="hidden sm:inline-block sm:align-middle sm:h-screen" aria-hidden="true">
                        &#8203;
                    </span>
                    <Transition.Child
                        as={Fragment}
                        enter="ease-out duration-300"
                        enterFrom="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"
                        enterTo="opacity-100 translate-y-0 sm:scale-100"
                        leave="ease-in duration-200"
                        leaveFrom="opacity-100 translate-y-0 sm:scale-100"
                        leaveTo="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"
                    >
                        <div className="inline-block align-bottom bg-white rounded-lg px-4 pt-5 pb-4 text-left overflow-hidden shadow-xl transform transition-all lg:my-8 sm:align-middle lg:max-w-lg lg:w-full lg:p-6">
                            <div className="hidden sm:block absolute top-0 right-0 pt-4 pr-4">
                                <button
                                    type="button"
                                    className="bg-white rounded-md text-gray-400 hover:text-gray-500 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
                                    onClick={props.onClose}
                                >
                                    <span className="sr-only">Close</span>
                                    <XIcon className="h-6 w-6" aria-hidden="true" />
                                </button>
                            </div>
                            <div className="sm:flex sm:items-start">
                                <div className="mx-auto flex-shrink-0 flex items-center justify-center h-12 w-12 rounded-full bg-indigo-100 sm:mx-0 sm:h-10 sm:w-10">
                                    <InformationCircleIcon className="h-6 w-6 text-indigo-600" aria-hidden="true" />
                                </div>
                                <div className="mt-3 text-center sm:mt-0 sm:ml-4 sm:text-left">
                                    <Dialog.Title as="h3" className="text-lg leading-6 font-medium text-gray-900">
                                        Install Package
                                    </Dialog.Title>
                                    <div className="mt-2">
                                        <label htmlFor="package" className="block text-sm font-medium text-gray-700">
                                            URL or Identifer
                                        </label>
                                
                                    <div className="mt-1 w-full">
                                        <input
                                            type="text"
                                            name="package"
                                            id="package"
                                            className="shadow-sm focus:ring-indigo-500 focus:border-indigo-500 block w-full sm:text-md border-gray-300 rounded-md"
                                            aria-describedby="package-description"
                                        />
                                    </div>
                                   
                                        <p className="mt-2 text-sm text-gray-500" id="package-description">
                                            e.g. https://github.com/aidmodels/detectron
                                        </p>
                                   </div>
                                </div>
                            </div>

                            <div className="mt-5">
                                <div className="w-full max-w-md p-2 mx-auto bg-white rounded-2xl">
                                    <Disclosure>
                                        {({ open }) => (
                                            <>
                                                <Disclosure.Button className="flex flex-row-reverse justify-between w-full px-4 py-2 text-sm font-medium text-right rounded-lg">
                                                    <ChevronUpIcon
                                                        className={`${open ? 'transform rotate-180' : ''
                                                            } w-5 h-5 text-purple-500`}
                                                    />
                                                </Disclosure.Button>
                                                <Disclosure.Panel className="px-4 pt-4 pb-2 text-sm text-gray-500 ">
                                                    <div>
                                                        <div>
                                                            Build Image
                                                            <Toggles />
                                                        </div>
                                                        <div>
                                                            Create Container
                                                            <Toggles />
                                                        </div>
                                                    </div>
                                                </Disclosure.Panel>
                                            </>
                                        )}
                                    </Disclosure>
                                </div>
                            </div>


                            <div className="mt-5 sm:mt-4 sm:flex sm:flex-row-reverse">
                                <button
                                    type="button"
                                    className="w-full inline-flex justify-center rounded-md border border-transparent shadow-sm px-4 py-2 bg-red-600 text-base font-medium text-white hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500 sm:ml-3 sm:w-auto sm:text-sm"
                                    onClick={props.onClose}
                                >
                                    Confirm
                                </button>
                                <button
                                    type="button"
                                    className="mt-3 w-full inline-flex justify-center rounded-md border border-gray-300 shadow-sm px-4 py-2 bg-white text-base font-medium text-gray-700 hover:text-gray-500 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 sm:mt-0 sm:w-auto sm:text-sm"
                                    onClick={props.onClose}
                                >
                                    Cancel
                                </button>
                            </div>
                        </div>
                    </Transition.Child>
                </div>
            </Dialog>
        </Transition.Root>
    )
}
