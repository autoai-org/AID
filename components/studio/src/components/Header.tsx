import { Fragment } from 'react'
import { Disclosure, Menu, Transition } from '@headlessui/react'
import {
    SearchIcon,
} from '@heroicons/react/solid'
import { MenuAlt1Icon, XIcon } from '@heroicons/react/outline'

const navigation = [
    { name: 'Dashboard', href: '#', current: true },
    { name: 'Domains', href: '#', current: false },
]
const userNavigation = [
    { name: 'Your Profile', href: '#' },
    { name: 'Settings', href: '#' },
    { name: 'Sign out', href: '#' },
]
function classNames(...classes: any) {
    return classes.filter(Boolean).join(' ')
}
export default function Header() {
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
                                        src="logo.png"
                                        alt="Workflow"
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
                                        {navigation.map((item) => (
                                            <a
                                                key={item.name}
                                                href={item.href}
                                                className="px-3 py-2 rounded-md text-sm font-medium text-indigo-200 hover:text-white"
                                                aria-current={item.current ? 'page' : undefined}
                                            >
                                                {item.name}
                                            </a>
                                        ))}
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
                                                            src="https://images.unsplash.com/photo-1517365830460-955ce3ccd263?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=256&h=256&q=80"
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

                    <Disclosure.Panel className="lg:hidden">
                        <div className="px-2 pt-2 pb-3 space-y-1">
                            {navigation.map((item) => (
                                <a
                                    key={item.name}
                                    href={item.href}
                                    className={classNames(
                                        item.current
                                            ? 'text-white bg-indigo-800'
                                            : 'text-indigo-200 hover:text-indigo-100 hover:bg-indigo-600',
                                        'block px-3 py-2 rounded-md text-base font-medium'
                                    )}
                                    aria-current={item.current ? 'page' : undefined}
                                >
                                    {item.name}
                                </a>
                            ))}
                        </div>
                        <div className="pt-4 pb-3 border-t border-indigo-800">
                            <div className="px-2 space-y-1">
                                {userNavigation.map((item) => (
                                    <a
                                        key={item.name}
                                        href={item.href}
                                        className="block px-3 py-2 rounded-md text-base font-medium text-indigo-200 hover:text-indigo-100 hover:bg-indigo-600"
                                    >
                                        {item.name}
                                    </a>
                                ))}
                            </div>
                        </div>
                    </Disclosure.Panel>
                </>
            )}
        </Disclosure>
    )
}