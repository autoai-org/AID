import { Popover } from '@headlessui/react'
import {
    MenuIcon,
} from '@heroicons/react/outline'
import { FirebaseAuthConsumer } from '@react-firebase/auth'
import firebase from "firebase/app";
import "firebase/auth";
import { useHistory } from "react-router-dom";

const navigation = [
    { name: 'Documents', href: 'https://aid.autoai.org' },
    { name: 'Status', href: 'https://status.autoai.org/' },
    { name: 'Donate', href: 'https://github.com/sponsors/xzyaoi' },
    { name: 'About', href: 'https://aid.autoai.org/docs/about/intro' },
]

export default function AIDFooter() {
    const history = useHistory();
    return (<Popover as="header" className="relative">
        {({ open }) => (
            <>
                <div className="bg-gray-900 pt-4 pb-4">
                    <nav
                        className="relative max-w-7xl mx-auto flex items-center justify-between px-4 sm:px-6"
                        aria-label="Global"
                    >
                        <div className="flex items-center flex-1">
                            <div className="flex items-center justify-between w-full md:w-auto">
                                <a href="/">
                                    <span className="sr-only">Workflow</span>
                                    <img
                                        className="h-8 w-auto sm:h-10"
                                        src="../logo.png"
                                        alt=""
                                    />
                                </a>
                                <div className="-mr-2 flex items-center md:hidden">
                                    <Popover.Button className="bg-gray-900 rounded-md p-2 inline-flex items-center justify-center text-gray-400 hover:bg-gray-800 focus:outline-none focus:ring-2 focus-ring-inset focus:ring-white">
                                        <span className="sr-only">Open main menu</span>
                                        <MenuIcon className="h-6 w-6" aria-hidden="true" />
                                    </Popover.Button>
                                </div>
                            </div>
                            <div className="hidden space-x-8 md:flex md:ml-10">
                                {navigation.map((item) => (
                                    <a
                                        key={item.name}
                                        href={item.href}
                                        className="text-base font-medium text-white hover:text-gray-300"
                                    >
                                        {item.name}
                                    </a>
                                ))}
                            </div>
                        </div>
                        <FirebaseAuthConsumer>
                            {({ isSignedIn, user, providerId }) => {
                                if (isSignedIn === true) {
                                    return (
                                        <div className="hidden md:flex md:items-center md:space-x-6">
                                            <a href="/profile" className="flex-shrink-0 group block">
                                                <div className="flex items-center">
                                                    <div>
                                                        <img
                                                            className="inline-block h-9 w-9 rounded-full"
                                                            src={user.photoURL}
                                                            alt=""
                                                        />
                                                    </div>
                                                    <div className="ml-3">
                                                        <p className="text-sm font-medium text-white group-hover:text-gray-900">{user.displayName}</p>
                                                    </div>
                                                </div>
                                            </a>
                                            <div
                                                onClick={()=>firebase.auth().signOut().then(function(res){
                                                    history.replace("/");
                                                })}
                                                className="inline-flex items-center px-4 py-2 border border-transparent text-base font-medium rounded-md text-white bg-gray-600 hover:bg-gray-700"
                                            >
                                                Sign out
                                            </div>
                                        </div>);
                                } else {
                                    return (
                                        <div className="hidden md:flex md:items-center md:space-x-6">
                                            <a href="/login" className="text-base font-medium text-white hover:text-gray-300">
                                                Log in
                                            </a>
                                            <a
                                                href="/signup"
                                                className="inline-flex items-center px-4 py-2 border border-transparent text-base font-medium rounded-md text-white bg-gray-600 hover:bg-gray-700"
                                            >
                                                Sign up
                                            </a>
                                        </div>
                                    )
                                }
                            }}
                        </FirebaseAuthConsumer>

                    </nav>
                </div>
            </>
        )}
    </Popover>

    )
}