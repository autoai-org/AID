import { useState } from 'react'
import { Switch } from '@headlessui/react'
import {
    BellIcon,
    CogIcon,
    CreditCardIcon,
    KeyIcon,
    UserCircleIcon,
    ViewGridAddIcon,
} from '@heroicons/react/outline'

const user = {
    name: 'Debbie Lewis',
    handle: 'deblewis',
    email: 'debbielewis@example.com',
    imageUrl:
        'https://images.unsplash.com/photo-1517365830460-955ce3ccd263?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=4&w=320&h=320&q=80',
}

const subNavigation = [
    { name: 'Profile', href: '#', icon: UserCircleIcon, current: true },
    { name: 'Account', href: '#', icon: CogIcon, current: false },
    { name: 'Password', href: '#', icon: KeyIcon, current: false },
    { name: 'Notifications', href: '#', icon: BellIcon, current: false },
    { name: 'Billing', href: '#', icon: CreditCardIcon, current: false },
    { name: 'Integrations', href: '#', icon: ViewGridAddIcon, current: false },
]

function classNames(...classes: any) {
    return classes.filter(Boolean).join(' ')
}

export default function Example() {
    const [availableToHire, setAvailableToHire] = useState(true)
    const [privateAccount, setPrivateAccount] = useState(false)
    const [allowCommenting, setAllowCommenting] = useState(true)
    const [allowMentions, setAllowMentions] = useState(true)

    return (
        <div>
            <div className="max-w-screen-xl mx-auto pb-6 px-4 sm:px-6 lg:pb-16 lg:px-8">
                <div className="bg-white rounded-lg shadow overflow-hidden">
                    <div className="divide-y divide-gray-200 lg:grid lg:grid-cols-12 lg:divide-y-0 lg:divide-x">
                        <aside className="py-6 lg:col-span-3">
                            <nav className="space-y-1">
                                {subNavigation.map((item) => (
                                    <a
                                        key={item.name}
                                        href={item.href}
                                        className={classNames(
                                            item.current
                                                ? 'bg-teal-50 border-teal-500 text-teal-700 hover:bg-teal-50 hover:text-teal-700'
                                                : 'border-transparent text-gray-900 hover:bg-gray-50 hover:text-gray-900',
                                            'group border-l-4 px-3 py-2 flex items-center text-sm font-medium'
                                        )}
                                        aria-current={item.current ? 'page' : undefined}
                                    >
                                        <item.icon
                                            className={classNames(
                                                item.current
                                                    ? 'text-teal-500 group-hover:text-teal-500'
                                                    : 'text-gray-400 group-hover:text-gray-500',
                                                'flex-shrink-0 -ml-1 mr-3 h-6 w-6'
                                            )}
                                            aria-hidden="true"
                                        />
                                        <span className="truncate">{item.name}</span>
                                    </a>
                                ))}
                            </nav>
                        </aside>

                        <form className="divide-y divide-gray-200 lg:col-span-9" action="#" method="POST">
                            {/* Profile section */}
                            <div className="py-6 px-4 sm:p-6 lg:pb-8">
                                <div>
                                    <h2 className="text-lg leading-6 font-medium text-gray-900">Profile</h2>
                                    <p className="mt-1 text-sm text-gray-500">
                                        This information will be displayed publicly so be careful what you share.
                                    </p>
                                </div>

                                <div className="mt-6 flex flex-col lg:flex-row">
                                    <div className="flex-grow space-y-6">
                                        <div>
                                            <label htmlFor="username" className="block text-sm font-medium text-gray-700">
                                                Username
                                            </label>
                                            <div className="mt-1 rounded-md shadow-sm flex">
                                                <span className="bg-gray-50 border border-r-0 border-gray-300 rounded-l-md px-3 inline-flex items-center text-gray-500 sm:text-sm">
                                                    workcation.com/
                                                </span>
                                                <input
                                                    type="text"
                                                    name="username"
                                                    id="username"
                                                    autoComplete="username"
                                                    className="focus:ring-sky-500 focus:border-sky-500 flex-grow block w-full min-w-0 rounded-none rounded-r-md sm:text-sm border-gray-300"
                                                    defaultValue={user.handle}
                                                />
                                            </div>
                                        </div>

                                        <div>
                                            <label htmlFor="about" className="block text-sm font-medium text-gray-700">
                                                About
                                            </label>
                                            <div className="mt-1">
                                                <textarea
                                                    id="about"
                                                    name="about"
                                                    rows={3}
                                                    className="shadow-sm focus:ring-sky-500 focus:border-sky-500 mt-1 block w-full sm:text-sm border border-gray-300 rounded-md"
                                                    defaultValue={''}
                                                />
                                            </div>
                                            <p className="mt-2 text-sm text-gray-500">
                                                Brief description for your profile. URLs are hyperlinked.
                                            </p>
                                        </div>
                                    </div>

                                    <div className="mt-6 flex-grow lg:mt-0 lg:ml-6 lg:flex-grow-0 lg:flex-shrink-0">
                                        <p className="text-sm font-medium text-gray-700" aria-hidden="true">
                                            Photo
                                        </p>
                                        <div className="mt-1 lg:hidden">
                                            <div className="flex items-center">
                                                <div
                                                    className="flex-shrink-0 inline-block rounded-full overflow-hidden h-12 w-12"
                                                    aria-hidden="true"
                                                >
                                                    <img className="rounded-full h-full w-full" src={user.imageUrl} alt="" />
                                                </div>
                                                <div className="ml-5 rounded-md shadow-sm">
                                                    <div className="group relative border border-gray-300 rounded-md py-2 px-3 flex items-center justify-center hover:bg-gray-50 focus-within:ring-2 focus-within:ring-offset-2 focus-within:ring-sky-500">
                                                        <label
                                                            htmlFor="user-photo"
                                                            className="relative text-sm leading-4 font-medium text-gray-700 pointer-events-none"
                                                        >
                                                            <span>Change</span>
                                                            <span className="sr-only"> user photo</span>
                                                        </label>
                                                        <input
                                                            id="user-photo"
                                                            name="user-photo"
                                                            type="file"
                                                            className="absolute w-full h-full opacity-0 cursor-pointer border-gray-300 rounded-md"
                                                        />
                                                    </div>
                                                </div>
                                            </div>
                                        </div>

                                        <div className="hidden relative rounded-full overflow-hidden lg:block">
                                            <img className="relative rounded-full w-40 h-40" src={user.imageUrl} alt="" />
                                            <label
                                                htmlFor="user-photo"
                                                className="absolute inset-0 w-full h-full bg-black bg-opacity-75 flex items-center justify-center text-sm font-medium text-white opacity-0 hover:opacity-100 focus-within:opacity-100"
                                            >
                                                <span>Change</span>
                                                <span className="sr-only"> user photo</span>
                                                <input
                                                    type="file"
                                                    id="user-photo"
                                                    name="user-photo"
                                                    className="absolute inset-0 w-full h-full opacity-0 cursor-pointer border-gray-300 rounded-md"
                                                />
                                            </label>
                                        </div>
                                    </div>
                                </div>

                                <div className="mt-6 grid grid-cols-12 gap-6">
                                    <div className="col-span-12 sm:col-span-6">
                                        <label htmlFor="first-name" className="block text-sm font-medium text-gray-700">
                                            First name
                                        </label>
                                        <input
                                            type="text"
                                            name="first-name"
                                            id="first-name"
                                            autoComplete="given-name"
                                            className="mt-1 block w-full border border-gray-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-sky-500 focus:border-sky-500 sm:text-sm"
                                        />
                                    </div>

                                    <div className="col-span-12 sm:col-span-6">
                                        <label htmlFor="last-name" className="block text-sm font-medium text-gray-700">
                                            Last name
                                        </label>
                                        <input
                                            type="text"
                                            name="last-name"
                                            id="last-name"
                                            autoComplete="family-name"
                                            className="mt-1 block w-full border border-gray-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-sky-500 focus:border-sky-500 sm:text-sm"
                                        />
                                    </div>

                                    <div className="col-span-12">
                                        <label htmlFor="url" className="block text-sm font-medium text-gray-700">
                                            URL
                                        </label>
                                        <input
                                            type="text"
                                            name="url"
                                            id="url"
                                            className="mt-1 block w-full border border-gray-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-sky-500 focus:border-sky-500 sm:text-sm"
                                        />
                                    </div>

                                    <div className="col-span-12 sm:col-span-6">
                                        <label htmlFor="company" className="block text-sm font-medium text-gray-700">
                                            Company
                                        </label>
                                        <input
                                            type="text"
                                            name="company"
                                            id="company"
                                            autoComplete="organization"
                                            className="mt-1 block w-full border border-gray-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-sky-500 focus:border-sky-500 sm:text-sm"
                                        />
                                    </div>
                                </div>
                            </div>

                            {/* Privacy section */}
                            <div className="pt-6 divide-y divide-gray-200">
                                <div className="px-4 sm:px-6">
                                    <div>
                                        <h2 className="text-lg leading-6 font-medium text-gray-900">Privacy</h2>
                                        <p className="mt-1 text-sm text-gray-500">
                                            Ornare eu a volutpat eget vulputate. Fringilla commodo amet.
                                        </p>
                                    </div>
                                    <ul className="mt-2 divide-y divide-gray-200">
                                        <Switch.Group as="li" className="py-4 flex items-center justify-between">
                                            <div className="flex flex-col">
                                                <Switch.Label as="p" className="text-sm font-medium text-gray-900" passive>
                                                    Available to hire
                                                </Switch.Label>
                                                <Switch.Description className="text-sm text-gray-500">
                                                    Nulla amet tempus sit accumsan. Aliquet turpis sed sit lacinia.
                                                </Switch.Description>
                                            </div>
                                            <Switch
                                                checked={availableToHire}
                                                onChange={setAvailableToHire}
                                                className={classNames(
                                                    availableToHire ? 'bg-teal-500' : 'bg-gray-200',
                                                    'ml-4 relative inline-flex flex-shrink-0 h-6 w-11 border-2 border-transparent rounded-full cursor-pointer transition-colors ease-in-out duration-200 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-sky-500'
                                                )}
                                            >
                                                <span
                                                    aria-hidden="true"
                                                    className={classNames(
                                                        availableToHire ? 'translate-x-5' : 'translate-x-0',
                                                        'inline-block h-5 w-5 rounded-full bg-white shadow transform ring-0 transition ease-in-out duration-200'
                                                    )}
                                                />
                                            </Switch>
                                        </Switch.Group>
                                        <Switch.Group as="li" className="py-4 flex items-center justify-between">
                                            <div className="flex flex-col">
                                                <Switch.Label as="p" className="text-sm font-medium text-gray-900" passive>
                                                    Make account private
                                                </Switch.Label>
                                                <Switch.Description className="text-sm text-gray-500">
                                                    Pharetra morbi dui mi mattis tellus sollicitudin cursus pharetra.
                                                </Switch.Description>
                                            </div>
                                            <Switch
                                                checked={privateAccount}
                                                onChange={setPrivateAccount}
                                                className={classNames(
                                                    privateAccount ? 'bg-teal-500' : 'bg-gray-200',
                                                    'ml-4 relative inline-flex flex-shrink-0 h-6 w-11 border-2 border-transparent rounded-full cursor-pointer transition-colors ease-in-out duration-200 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-sky-500'
                                                )}
                                            >
                                                <span
                                                    aria-hidden="true"
                                                    className={classNames(
                                                        privateAccount ? 'translate-x-5' : 'translate-x-0',
                                                        'inline-block h-5 w-5 rounded-full bg-white shadow transform ring-0 transition ease-in-out duration-200'
                                                    )}
                                                />
                                            </Switch>
                                        </Switch.Group>
                                        <Switch.Group as="li" className="py-4 flex items-center justify-between">
                                            <div className="flex flex-col">
                                                <Switch.Label as="p" className="text-sm font-medium text-gray-900" passive>
                                                    Allow commenting
                                                </Switch.Label>
                                                <Switch.Description className="text-sm text-gray-500">
                                                    Integer amet, nunc hendrerit adipiscing nam. Elementum ame
                                                </Switch.Description>
                                            </div>
                                            <Switch
                                                checked={allowCommenting}
                                                onChange={setAllowCommenting}
                                                className={classNames(
                                                    allowCommenting ? 'bg-teal-500' : 'bg-gray-200',
                                                    'ml-4 relative inline-flex flex-shrink-0 h-6 w-11 border-2 border-transparent rounded-full cursor-pointer transition-colors ease-in-out duration-200 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-sky-500'
                                                )}
                                            >
                                                <span
                                                    aria-hidden="true"
                                                    className={classNames(
                                                        allowCommenting ? 'translate-x-5' : 'translate-x-0',
                                                        'inline-block h-5 w-5 rounded-full bg-white shadow transform ring-0 transition ease-in-out duration-200'
                                                    )}
                                                />
                                            </Switch>
                                        </Switch.Group>
                                        <Switch.Group as="li" className="py-4 flex items-center justify-between">
                                            <div className="flex flex-col">
                                                <Switch.Label as="p" className="text-sm font-medium text-gray-900" passive>
                                                    Allow mentions
                                                </Switch.Label>
                                                <Switch.Description className="text-sm text-gray-500">
                                                    Adipiscing est venenatis enim molestie commodo eu gravid
                                                </Switch.Description>
                                            </div>
                                            <Switch
                                                checked={allowMentions}
                                                onChange={setAllowMentions}
                                                className={classNames(
                                                    allowMentions ? 'bg-teal-500' : 'bg-gray-200',
                                                    'ml-4 relative inline-flex flex-shrink-0 h-6 w-11 border-2 border-transparent rounded-full cursor-pointer transition-colors ease-in-out duration-200 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-sky-500'
                                                )}
                                            >
                                                <span
                                                    aria-hidden="true"
                                                    className={classNames(
                                                        allowMentions ? 'translate-x-5' : 'translate-x-0',
                                                        'inline-block h-5 w-5 rounded-full bg-white shadow transform ring-0 transition ease-in-out duration-200'
                                                    )}
                                                />
                                            </Switch>
                                        </Switch.Group>
                                    </ul>
                                </div>
                                <div className="mt-4 py-4 px-4 flex justify-end sm:px-6">
                                    <button
                                        type="button"
                                        className="bg-white border border-gray-300 rounded-md shadow-sm py-2 px-4 inline-flex justify-center text-sm font-medium text-gray-700 hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-sky-500"
                                    >
                                        Cancel
                                    </button>
                                    <button
                                        type="submit"
                                        className="ml-5 bg-sky-700 border border-transparent rounded-md shadow-sm py-2 px-4 inline-flex justify-center text-sm font-medium text-white hover:bg-sky-800 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-sky-500"
                                    >
                                        Save
                                    </button>
                                </div>
                            </div>
                        </form>
                    </div>
                </div>
            </div>
        </div>
    )
}