import { Fragment, useEffect, useState } from "react";
import { Dialog, Transition, Disclosure } from '@headlessui/react'
import { MailIcon, PhoneIcon } from '@heroicons/react/solid'
import { ALL_IMAGES } from "../../services/apis/queries";
import { restclient } from "../../services/apis";
import { useDispatch } from "react-redux";
import { setIsLoading } from "../../services/store/connectivity/server";
const people = [
    {
        name: 'Jane Cooper',
        title: 'Regional Paradigm Technician',
        role: 'Admin',
        email: 'janecooper@example.com',
        telephone: '+1-202-555-0170',
        imageUrl:
            'https://images.unsplash.com/photo-1494790108377-be9c29b29330?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=4&w=256&h=256&q=60',
    },
    {
        name: 'Jane Cooper',
        title: 'Regional Paradigm Technician',
        role: 'Admin',
        email: 'janecooper@example.com',
        telephone: '+1-202-555-0170',
        imageUrl:
            'https://images.unsplash.com/photo-1494790108377-be9c29b29330?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=4&w=256&h=256&q=60',
    },
    {
        name: 'Jane Cooper',
        title: 'Regional Paradigm Technician',
        role: 'Admin',
        email: 'janecooper@example.com',
        telephone: '+1-202-555-0170',
        imageUrl:
            'https://images.unsplash.com/photo-1494790108377-be9c29b29330?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=4&w=256&h=256&q=60',
    },
    {
        name: 'Jane Cooper',
        title: 'Regional Paradigm Technician',
        role: 'Admin',
        email: 'janecooper@example.com',
        telephone: '+1-202-555-0170',
        imageUrl:
            'https://images.unsplash.com/photo-1494790108377-be9c29b29330?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=4&w=256&h=256&q=60',
    },
    // More people...
]

export default function ImageList(props: any) {
    const [images, setImages] = useState<object[]>([]);
    const dispatch = useDispatch();
    useEffect(() => {
        dispatch(setIsLoading(true))
        restclient.query(ALL_IMAGES).then((res:any) => {
            console.log(res.data.data.images)
            setImages(res.data.data.images)
            dispatch(setIsLoading(false))
        })
    }, [])
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
                        <div className="inline-block align-bottom bg-white rounded-lg px-4 pt-5 pb-4 text-left overflow-hidden shadow-xl transform transition-all lg:my-8 lg:align-middle lg:max-w-2xl lg:w-full lg:p-6">
                            <div className="w-full max-w-2xl mx-auto bg-white rounded-2xl">
                                <Disclosure>
                                    {({ open }) => (
                                        <>
                                            <ul role="list" className="grid grid-cols-1 gap-6 sm:grid-cols-2 lg:grid-cols-3">
                                                {images.map((image:any) => (
                                                    <li key={image.title} className="col-span-1 bg-white rounded-lg shadow divide-y divide-gray-200">
                                                        <div className="w-full flex items-center justify-between p-6 space-x-6">
                                                            <div className="flex-1 truncate">
                                                                <div className="flex items-center space-x-3">
                                                                    <h3 className="text-gray-900 text-sm font-medium truncate">{image.title}</h3>
                                                                    <span className="flex-shrink-0 inline-block px-2 py-0.5 text-green-800 text-xs font-medium bg-green-100 rounded-full">
                                                                        {image.title}
                                                                    </span>
                                                                </div>
                                                                <p className="mt-1 text-gray-500 text-sm truncate">{image.title}</p>
                                                            </div>
                                                        </div>
                                                        <div>
                                                            <div className="-mt-px flex divide-x divide-gray-200">
                                                                
                                                                <div className="-ml-px w-0 flex-1 flex">
                                                                    <div
                                                                        className="relative w-0 flex-1 inline-flex items-center justify-center py-4 text-sm text-gray-700 font-medium border border-transparent rounded-br-lg hover:text-gray-500"
                                                                    >
                                                                        
                                                                        <span className="ml-3">Create</span>
                                                                    </div>
                                                                </div>
                                                            </div>
                                                        </div>
                                                    </li>
                                                ))}
                                            </ul>
                                        </>
                                    )}
                                </Disclosure>
                            </div>

                        </div>
                    </Transition.Child>
                </div>
            </Dialog>
        </Transition.Root>
    )
}