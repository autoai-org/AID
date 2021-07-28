import { Fragment } from 'react'
import { Disclosure, Menu, Transition } from '@headlessui/react'
import {
    SearchIcon,
} from '@heroicons/react/solid'
import { MenuAlt1Icon, XIcon } from '@heroicons/react/outline'
import { serverEndpoint, setServerEndpoint, getServerEndpoint } from '../services/apis'
import { store } from '../services/store/store'
const userNavigation = [
    { name: 'Your Profile', href: '#' },
    { name: 'Settings', href: '#' },
    { name: 'Log Out', href: '#' },
]

function classNames(...classes: any) {
    return classes.filter(Boolean).join(' ')
}
export default function Header() {
    function disconnect() {
        setServerEndpoint("")
    }
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
                                            id="search"
                                            name="search"
                                            className="block w-full pl-10 pr-3 py-2 border border-transparent rounded-md leading-5 bg-indigo-400 bg-opacity-25 text-indigo-100 placeholder-indigo-200 focus:outline-none focus:bg-white focus:ring-0 focus:placeholder-gray-400 focus:text-gray-900 sm:text-sm"
                                            placeholder="Search projects"
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

                                        <Menu as="div" className="ml-4 relative flex-shrink-0">
                                            {({ open }) => (
                                                <>
                                                    <div>
                                                        <Menu.Button className="flex text-sm rounded-full hover:text-white text-indigo-200">
                                                            Connnected
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
                                    {/* Profile dropdown */}
                                    <Menu as="div" className="ml-4 relative flex-shrink-0">
                                        {({ open }) => (
                                            <>
                                                <div>
                                                    <Menu.Button className="bg-indigo-700 flex text-sm rounded-full text-white focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-offset-indigo-700 focus:ring-white">
                                                        <span className="sr-only">Open user menu</span>
                                                        <img
                                                            className="h-8 w-8 rounded-full"
                                                            src="https://xzyaoi.github.io/avatar.jpg"
                                                            alt=""
                                                        />
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
                                                        {userNavigation.map((item) => (
                                                            <Menu.Item key={item.name}>
                                                                {({ active }) => (
                                                                    <a
                                                                        href={item.href}
                                                                        className={classNames(
                                                                            active ? 'bg-gray-100' : '',
                                                                            'block px-4 py-2 text-sm text-gray-700'
                                                                        )}
                                                                    >

                                                                        {item.name}
                                                                    </a>
                                                                )}
                                                            </Menu.Item>
                                                        ))}
                                                    </Menu.Items>
                                                </Transition>
                                            </>
                                        )}
                                    </Menu>
                                </div>
                            </div>
                        </div>
                    </div>
                </>
            )}
        </Disclosure>
    )
}