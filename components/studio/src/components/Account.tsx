import { useState } from 'react'
import {
    BadgeCheckIcon,
    CollectionIcon,
} from '@heroicons/react/solid'
import { FirebaseAuthConsumer } from '@react-firebase/auth'
import InstallInstallPackageDialog from "./Workflow/InstallPackageDialog"
import ImageList from './Workflow/ImageList'
import "firebase/auth";

export default function Account() {
    const [open, setOpen] = useState(false)
    const [openRepos, setOpenRepos] = useState(false)
    const openInstallDialog = () => {
        setOpen(true)
    }
    const openRepoList = () =>{
        setOpenRepos(true)
    }
    const CloseInstallDialog = () => {
        setOpen(false)
    }
    const CloseRepoList = () => {
        setOpenRepos(false)
    }
    return (
        <div className="xl:flex-shrink-0 xl:w-64 xl:border-r xl:border-gray-200 bg-white">
            <div className="pl-4 pr-6 py-6 sm:pl-6 lg:pl-8 xl:pl-0">
                <div className="flex items-center justify-between">
                    <div className="flex-1 space-y-8">
                        <div className="space-y-8 sm:space-y-0 sm:flex sm:justify-between sm:items-center xl:block xl:space-y-8">
                            {/* Profile */}
                            <FirebaseAuthConsumer>
                                {({ isSignedIn, user, providerId }) => {
                                    if (isSignedIn) {
                                        return (
                                            <div className="flex items-center space-x-3">
                                                <div className="flex-shrink-0 h-12 w-12">
                                                    <img
                                                        className="h-12 w-12 rounded-full"
                                                        src={user.photoURL}
                                                        alt=""
                                                    />
                                                </div>

                                                <div className="space-y-1">
                                                    <div className="text-sm font-medium text-gray-900">{user.displayName}</div>
                                                </div>
                                            </div>
                                        )
                                    } else {
                                        return (
                                            <div className="flex flex-col sm:flex-row xl:flex-col">
                                                <a
                                                    href="/signin"
                                                    type="button"
                                                    className="inline-flex items-center justify-center px-4 py-2 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 xl:w-full"
                                                >
                                                    Login
                                                </a>
                                            </div>
                                        )

                                    }
                                }}

                            </FirebaseAuthConsumer>
                            {/* Action buttons */}
                            <div className="flex flex-col sm:flex-row xl:flex-col">
                                <button
                                    onClick={openInstallDialog}
                                    type="button"
                                    className="inline-flex items-center justify-center px-4 py-2 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 xl:w-full"
                                >
                                    Install Packages
                                </button>
                                <button
                                    type="button"
                                    onClick={openRepoList}
                                    className="mt-3 inline-flex items-center justify-center px-4 py-2 border border-gray-300 shadow-sm text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 sm:mt-0 sm:ml-3 xl:ml-0 xl:mt-3 xl:w-full"
                                >
                                    Installed Repositories
                                </button>
                            </div>
                        </div>
                        {/* Meta info */}
                        <div className="flex flex-col space-y-6 sm:flex-row sm:space-y-0 sm:space-x-8 xl:flex-col xl:space-x-0 xl:space-y-6">
                            <div className="flex items-center space-x-2">
                                <BadgeCheckIcon className="h-5 w-5 text-gray-400" aria-hidden="true" />
                                <span className="text-sm text-gray-500 font-medium">Early Adopter</span>
                            </div>
                            <div className="flex items-center space-x-2">
                                <CollectionIcon className="h-5 w-5 text-gray-400" aria-hidden="true" />
                                <span className="text-sm text-gray-500 font-medium">Version 1.3.0</span>
                            </div>

                        </div>
                    </div>
                </div>
            </div>
            <InstallInstallPackageDialog open={open} onClose={CloseInstallDialog}/>
            <ImageList open={openRepos} onClose={CloseRepoList}/>
        </div>
        
    )
}