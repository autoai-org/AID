import {
    UserCircleIcon,
    ViewGridAddIcon,
} from '@heroicons/react/outline'
import { FirebaseAuthConsumer } from '@react-firebase/auth'

const subNavigation = [
    { name: 'Profile', href: '/profile', icon: UserCircleIcon, current: true },
    { name: 'New Models', href: '/publish', icon: ViewGridAddIcon, current: false },
]

function classNames(...classes: any) {
    return classes.filter(Boolean).join(' ')
}

export default function Example() {
    return (
        <div>
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
                    <div className="divide-y divide-gray-200 lg:col-span-9">
                        <FirebaseAuthConsumer>
                            {({ isSignedIn, user, providerId }) => {
                                console.log(user)
                                if (isSignedIn === true) {
                                    return (
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
                                                                https://hub.autoai.org/users/
                                                            </span>
                                                            <input
                                                                type="text"
                                                                name="username"
                                                                id="username"
                                                                autoComplete="username"
                                                                className="focus:ring-sky-500 focus:border-sky-500 flex-grow block w-full min-w-0 rounded-none rounded-r-md sm:text-sm border-gray-300"
                                                                defaultValue={user.uid}
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
                                                                <img className="rounded-full h-full w-full" src={user.photoURL} alt="" />
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
                                                        <img className="relative rounded-full w-40 h-40" src={user.photoURL} alt="" />
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
                                    )
                                }
                            }}
                        </FirebaseAuthConsumer>
                    </div>
                </div>
            </div>
        </div>
    )
}