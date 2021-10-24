import { useState } from "react"
import firebase from "firebase/app";
import "firebase/auth";
import { Octokit } from "@octokit/rest";
import ReactMarkdown from 'react-markdown'
import toml from 'toml'
import { postNewRepo } from '../../services/api'
import { CheckCircleIcon } from '@heroicons/react/solid'

let octokit = new Octokit();
export default function PublishForm() {

    const githubAuthProvider = new firebase.auth.GithubAuthProvider();
    
    const [loggedInFromGithub, setLoggedInFromGitHub] = useState(false);
    const [repositories, setRepositories] = useState<object[]>([]);
    const [logins, setLogins] = useState<string[]>([])
    const [selectedLogin, setSelectedLogin] = useState("")
    const [selectedRepo, setSelectedRepo] = useState("")
    const [repoConfirmed, setRepoConfirmed] = useState(false)
    const [readme, setReadme] = useState("")
    const [aidTOML, setAIDToml] = useState<any>({})
    const [added, setAdded] = useState(false)

    function GitHubImport() {
        let usernames: string[] = []
        let repositories: object[] = []
        firebase.auth().signInWithPopup(githubAuthProvider).then(function (res: any) {
            usernames.push(res.additionalUserInfo.username)
            let accessToken = res.credential.accessToken
            octokit = new Octokit({
                auth: accessToken,
            });

            octokit.rest.orgs.listForAuthenticatedUser().then(function (res) {
                res.data.map(function (each: any) {
                    usernames.push(each.login)
                    octokit.rest.repos.listForOrg({ org: each.login }).then(function (res) {
                        res.data.map(function (eachRepo: any) {
                            repositories.push({
                                login: each.login,
                                repo: eachRepo.name
                            })
                            return true
                        })
                        octokit.rest.repos.listForAuthenticatedUser().then(function (res) {
                            res.data.map(function (each: any) {
                                repositories.push({
                                    login: usernames[0],
                                    repo: each.name
                                })
                                return true
                            }
                            )
                            setLogins(usernames)
                            setSelectedLogin(usernames[0])
                            setRepositories(repositories)
                            setLoggedInFromGitHub(true)
                        })
                    })
                    return true
                })

            })
        })
    }
    function loadReadme() {
        octokit.rest.repos.getReadme({
            owner: selectedLogin,
            repo: selectedRepo,
        }).then(function (res: any) {
            setReadme(atob(res.data.content))
            octokit.rest.repos.getContent({
                owner: selectedLogin,
                repo: selectedRepo,
                path: 'aid.toml',
            }).then(function (res: any) {

                setAIDToml(toml.parse(atob(res.data.content)))
                setRepoConfirmed(true)
            })
        })
    }
    function submit() {
        postNewRepo({
            'name': selectedRepo,
            'vendor': selectedLogin,
            'description': aidTOML.package.tagline,
            'githubURL': "https://github.com/" + selectedLogin + "/" + selectedRepo
        }).then(function (res) {
            setAdded(true)
        })
    }
    return (
        <form className="space-y-8 divide-y divide-gray-200">
            <div className="space-y-8 divide-y divide-gray-200">
                {added &&
                    <div className="rounded-md bg-green-50 p-4">
                        <div className="flex">
                            <div className="flex-shrink-0">
                                <CheckCircleIcon className="h-5 w-5 text-green-400" aria-hidden="true" />
                            </div>
                            <div className="ml-3">
                                <h3 className="text-sm font-medium text-green-800">Submission Completed</h3>
                                <div className="mt-2 text-sm text-green-700">
                                    <p>Your model have been added!</p>
                                </div>
                            </div>
                        </div>
                    </div>
                }
                <div>
                    <div>
                        <h3 className="text-lg leading-6 font-medium text-gray-900">Publish New Models</h3>
                        <p className="mt-1 text-sm text-gray-500">
                            This information will be displayed publicly so be careful what you share.
                        </p>
                    </div>

                    <div className="mt-6 grid grid-cols-1 gap-y-6 gap-x-4 sm:grid-cols-6">
                        <div className="sm:col-span-4">
                            <button
                                onClick={() => {
                                    GitHubImport()
                                }}
                                type="button"
                                className="ml-3 inline-flex justify-center py-2 px-4 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
                            >
                                Import from GitHub
                            </button>
                        </div>
                    </div>
                    {loggedInFromGithub &&
                        <div className="mt-6 grid grid-cols-4 gap-y-6 gap-x-4 sm:grid-cols-6">
                            <div className="sm:col-span-1">
                                <label htmlFor="logins" className="block text-sm font-medium text-gray-700 sm:mt-px sm:pt-2">
                                    GitHub Users/Organizations
                                </label>
                                <div className="mt-1 sm:mt-0 sm:col-span-2">
                                    <select
                                        onChange={(e) => { setSelectedLogin(e.target.value) }}
                                        id="logins"
                                        name="country"
                                        autoComplete="country"
                                        className="max-w-lg block focus:ring-indigo-500 focus:border-indigo-500 w-full shadow-sm sm:max-w-xs sm:text-sm border-gray-300 rounded-md"
                                    >
                                        {logins.map(function (username, i) {
                                            return <option key={i}>{username}</option>;
                                        })}
                                    </select>
                                </div>
                            </div>
                            <div className="sm:col-span-2">
                                <label htmlFor="logins" className="block text-sm font-medium text-gray-700 sm:mt-px sm:pt-2">
                                    Model Name
                                </label>
                                <div className="mt-1 sm:mt-0 sm:col-span-2">
                                    <select
                                        id="logins"
                                        name="country"
                                        onChange={(e) => { setSelectedRepo(e.target.value) }}
                                        autoComplete="country"
                                        className="max-w-lg block focus:ring-indigo-500 focus:border-indigo-500 w-full shadow-sm sm:max-w-xs sm:text-sm border-gray-300 rounded-md"
                                    >
                                        {repositories.map(function (each: any, i) {
                                            if (each.login === selectedLogin) {
                                                return <option key={i}>{each.repo}</option>;
                                            }
                                            return null;

                                        })}
                                    </select>
                                </div>
                            </div>
                            <div className="sm:col-span-1">
                                <label htmlFor="logins" className="block text-sm font-medium text-gray-700 sm:mt-px sm:pt-2">
                                    &nbsp;
                                </label>
                                <div className="mt-1 sm:mt-0 sm:col-span-1">
                                    <button type="button"
                                        onClick={loadReadme}
                                        className="ml-3 inline-flex justify-center py-2 px-4 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">Confirm</button>
                                </div>
                            </div>
                        </div>}

                    {repoConfirmed &&
                        <div className="mt-6 grid grid-cols-4 gap-y-6 gap-x-4 sm:grid-cols-6">
                            <div className="sm:col-span-6">
                                <label htmlFor="about" className="block text-sm font-medium text-gray-700">
                                    Tagline
                                </label>
                                <input
                                    type="text"
                                    name="tagline"
                                    id="tagline"
                                    disabled
                                    className="shadow-sm focus:ring-indigo-500 focus:border-indigo-500 block w-full sm:text-sm border-gray-300 rounded-md"
                                    value={aidTOML.package.tagline}
                                />

                            </div>
                            <div className="sm:col-span-6">
                                <label htmlFor="about" className="block text-sm font-medium text-gray-700">
                                    About
                                </label>
                                <div className="mt-1 prose">
                                    <ReactMarkdown>{readme}</ReactMarkdown>
                                </div>

                            </div>
                        </div>
                    }

                </div>

            </div>

            <div className="pt-5">
                <div className="flex justify-end">
                    <button
                        type="button"
                        className="bg-white py-2 px-4 border border-gray-300 rounded-md shadow-sm text-sm font-medium text-gray-700 hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
                    >
                        Cancel
                    </button>
                    <button
                        onClick={submit}
                        type="button"
                        className="ml-3 inline-flex justify-center py-2 px-4 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
                    >
                        Save
                    </button>
                </div>
            </div>
        </form>
    )
}