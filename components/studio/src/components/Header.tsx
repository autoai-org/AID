import { Fragment, useState } from 'react'
import { Disclosure, Menu, Transition } from '@headlessui/react'
import {
    SearchIcon,
} from '@heroicons/react/solid'
import { MenuAlt1Icon, XIcon } from '@heroicons/react/outline'
import { setServerEndpoint, getServerEndpoint } from '../services/apis'
import Terminal from './Common/Terminal'
function classNames(...classes: any) {
    return classes.filter(Boolean).join(' ')
}
export default function Header() {
    function disconnect() {
        setServerEndpoint("")
    }
    const [termOpen, setOpen] = useState(false)
    const handleKeyDown = (event: any) => {
        if (event.key === 'Enter') {
            if (queryKW.trim() !== '') {
                window.open('https://hub.autoai.dev/search?kw=' + queryKW.trim())
            }
        }
    }
    const openTerminalDialog = () => {
        console.log('===')
        setOpen(true)
    }
    const CloseTerminalDialog = () => {
        setOpen(false)
    }
    const [queryKW, setQueryKW] = useState("");
    return (
        <Disclosure as="nav" className="flex-shrink-0 bg-indigo-600">
            {({ open }) => (
                <>
                    <div className="max-w-7xl mx-auto px-2 sm:px-4 lg:px-8">
                        <div className="relative flex items-center justify-between h-16">
                            {/* Logo section */}
                            <div className="flex items-center px-2 lg:px-0 xl:w-64">
                                <div className="flex-shrink-0">
                                    <img
                                        className="h-8 w-auto"
                                        src="/logo.png"
                                        alt="AID Project"
                                    />
                                </div>
                            </div>

                            {/* Search section */}
                            <div className="flex-1 flex justify-center lg:justify-end">
                                <div className="w-full px-2 lg:px-6">
                                    <label htmlFor="search" className="sr-only">
                                        Search projects
                                    </label>
                                    <div className="relative text-indigo-200 focus-within:text-gray-400">
                                        <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                                            <SearchIcon className="h-5 w-5" aria-hidden="true" />
                                        </div>
                                        <input
                                            onKeyDown={handleKeyDown}
                                            id="search"
                                            name="search"
                                            onChange={e => {
                                                setQueryKW(e.target.value)
                                            }}
                                            className="block w-full pl-10 pr-3 py-2 border border-transparent rounded-md leading-5 bg-indigo-400 bg-opacity-25 text-indigo-100 placeholder-indigo-200 focus:outline-none focus:bg-white focus:ring-0 focus:placeholder-gray-400 focus:text-gray-900 sm:text-sm"
                                            placeholder="Search packages"
                                            type="search"
                                        />
                                    </div>
                                </div>
                            </div>
                            <div className="flex lg:hidden">
                                {/* Mobile menu button */}
                                <Disclosure.Button className="bg-indigo-600 inline-flex items-center justify-center p-2 rounded-md text-indigo-400 hover:text-white hover:bg-indigo-600 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-offset-indigo-600 focus:ring-white">
                                    <span className="sr-only">Open main menu</span>
                                    {open ? (
                                        <XIcon className="block h-6 w-6" aria-hidden="true" />
                                    ) : (
                                        <MenuAlt1Icon className="block h-6 w-6" aria-hidden="true" />
                                    )}
                                </Disclosure.Button>
                            </div>
                            {/* Links section */}
                            <div className="hidden lg:block lg:w-80">
                                <div className="flex items-center justify-end">
                                    <div className="flex">
                                        <div className="ml-4 relative flex-shrink-0">
                                            <button className="flex text-sm rounded-full hover:text-white text-indigo-100" onClick={()=>{alert("Not Available Yet.")}}>
                                                <svg xmlns="http://www.w3.org/2000/svg" className="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                                                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M8 9l3 3-3 3m5 0h3M5 20h14a2 2 0 002-2V6a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
                                                    
                                                </svg>
                                            </button>
                                        </div>
                                        <Menu as="div" className="ml-4 relative flex-shrink-0">
                                            {({ open }) => (
                                                <>
                                                    <div>
                                                        <Menu.Button className="flex text-sm rounded-full hover:text-white text-indigo-100">
                                                            Connnected to {getServerEndpoint()}
                                                        </Menu.Button>
                                                    </div>
                                                    <Transition
                                                        show={open}
                                                        as={Fragment}
                                                        enter="transition ease-out duration-100"
                                                        enterFrom="transform opacity-0 scale-95"
                                                        enterTo="transform opacity-100 scale-100"
                                                        leave="transition ease-in duration-75"
                                                        leaveFrom="transform opacity-100 scale-100"
                                                        leaveTo="transform opacity-0 scale-95"
                                                    >
                                                        <Menu.Items
                                                            static
                                                            className="origin-top-right absolute right-0 mt-2 w-48 rounded-md shadow-lg py-1 bg-white ring-1 ring-black ring-opacity-5 focus:outline-none"
                                                        >
                                                            <div>
                                                                <Menu.Item>
                                                                    {({ active }) => (
                                                                        <button
                                                                            onClick={disconnect}
                                                                            className={classNames(
                                                                                'text-gray-700 flex justify-between px-4 py-2 text-sm'
                                                                            )}
                                                                        >
                                                                            <span>Disconnect</span>
                                                                        </button>
                                                                    )}
                                                                </Menu.Item>
                                                            </div>

                                                        </Menu.Items>
                                                    </Transition>
                                                </>
                                            )}
                                        </Menu>
                                    </div>
                                    
                                    <Terminal open={termOpen} onClose={CloseTerminalDialog}/>
                                </div>
                            </div>
                        </div>
                    </div>
                </>
            )}
            
        </Disclosure>
    )
}