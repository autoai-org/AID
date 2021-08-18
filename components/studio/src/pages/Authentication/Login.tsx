import firebase from "firebase/app";
import "firebase/auth";
import { useState } from "react";
import Alert from '../../components/Common/Alert';
import { useHistory } from "react-router-dom";
export default function Login() {
    const [email, setEmail] = useState("")
    const [password, setPassword] = useState("")
    const [error, setError] = useState("")
    const googleAuthProvider = new firebase.auth.GoogleAuthProvider();
    const githubAuthProvider = new firebase.auth.GithubAuthProvider();
    const history = useHistory();

    function handleEmailChange(e: any) {
        setEmail(e.target.value)
    }
    function handlePasswordChange(e: any) {
        setPassword(e.target.value)
    }
    return (
        <div className="relative bg-gray-800 overflow-hidden" style={{ height: '100vh' }}>
            <div className="hidden sm:block sm:absolute sm:inset-0" aria-hidden="true">
                <svg
                    className="absolute bottom-0 right-0 transform translate-x-1/2 mb-48 text-gray-700 lg:top-0 lg:mt-28 lg:mb-0 xl:transform-none xl:translate-x-0"
                    width={364}
                    height={384}
                    viewBox="0 0 364 384"
                    fill="none"
                >
                    <defs>
                        <pattern
                            id="eab71dd9-9d7a-47bd-8044-256344ee00d0"
                            x={0}
                            y={0}
                            width={20}
                            height={20}
                            patternUnits="userSpaceOnUse"
                        >
                            <rect x={0} y={0} width={4} height={4} fill="currentColor" />
                        </pattern>
                    </defs>
                    <rect width={364} height={384} fill="url(#eab71dd9-9d7a-47bd-8044-256344ee00d0)" />
                </svg>
            </div>
            <div className="relative pt-6 pb-16 sm:pb-24">
                <main className="mt-16 sm:mt-24">
                    <div className="mx-auto max-w-7xl">
                        
                        <div className="mt-16 sm:mt-24 lg:mt-0 lg:col-span-6">
                        {error !== "" &&
                            <Alert title={"An error occurred during the request"} message={error}/>
                        }
                            <div className="bg-white sm:max-w-md sm:w-full sm:mx-auto sm:rounded-lg sm:overflow-hidden">
                                <div className="px-4 py-8 sm:px-10">
                                    <div>
                                        <p className="text-sm font-medium text-gray-700">Sign in with</p>
                                        <div className="mt-1 grid grid-cols-3 gap-3">
                                            <div>
                                                <div
                                                    onClick={() => firebase.auth().signInWithPopup(googleAuthProvider).then(function(res){
                                                        history.replace("/");
                                                    }).catch(function(err){
                                                        setError(err.message)
                                                    })}
                                                    className="w-full inline-flex justify-center py-2 px-4 border border-gray-300 rounded-md shadow-sm bg-white text-sm font-medium text-gray-500 hover:bg-gray-50"
                                                >
                                                    <span className="sr-only">Sign in with Google</span>
                                                    <svg className="w-5 h-5" aria-hidden="true" fill="currentColor" viewBox="0 0 24 24">
                                                        <g transform="matrix(1, 0, 0, 1, 27.009001, -39.238998)">
                                                            <path fill="#4285F4" d="M -3.264 51.509 C -3.264 50.719 -3.334 49.969 -3.454 49.239 L -14.754 49.239 L -14.754 53.749 L -8.284 53.749 C -8.574 55.229 -9.424 56.479 -10.684 57.329 L -10.684 60.329 L -6.824 60.329 C -4.564 58.239 -3.264 55.159 -3.264 51.509 Z" />
                                                            <path fill="#34A853" d="M -14.754 63.239 C -11.514 63.239 -8.804 62.159 -6.824 60.329 L -10.684 57.329 C -11.764 58.049 -13.134 58.489 -14.754 58.489 C -17.884 58.489 -20.534 56.379 -21.484 53.529 L -25.464 53.529 L -25.464 56.619 C -23.494 60.539 -19.444 63.239 -14.754 63.239 Z" />
                                                            <path fill="#FBBC05" d="M -21.484 53.529 C -21.734 52.809 -21.864 52.039 -21.864 51.239 C -21.864 50.439 -21.724 49.669 -21.484 48.949 L -21.484 45.859 L -25.464 45.859 C -26.284 47.479 -26.754 49.299 -26.754 51.239 C -26.754 53.179 -26.284 54.999 -25.464 56.619 L -21.484 53.529 Z" />
                                                            <path fill="#EA4335" d="M -14.754 43.989 C -12.984 43.989 -11.404 44.599 -10.154 45.789 L -6.734 42.369 C -8.804 40.429 -11.514 39.239 -14.754 39.239 C -19.444 39.239 -23.494 41.939 -25.464 45.859 L -21.484 48.949 C -20.534 46.099 -17.884 43.989 -14.754 43.989 Z" />
                                                        </g>
                                                    </svg>
                                                </div>
                                            </div>

                                            <div>
                                                <div
                                                    onClick={() => firebase.auth().signInWithPopup(githubAuthProvider).then(function(res){
                                                        history.replace("/");
                                                    }).catch(function(err){
                                                        setError(err.message)
                                                    })}
                                                    className="w-full inline-flex justify-center py-2 px-4 border border-gray-300 rounded-md shadow-sm bg-white text-sm font-medium text-gray-500 hover:bg-gray-50"
                                                >
                                                    <span className="sr-only">Sign in with GitHub</span>
                                                    <svg className="w-5 h-5" aria-hidden="true" fill="currentColor" viewBox="0 0 20 20">
                                                        <path
                                                            fillRule="evenodd"
                                                            d="M10 0C4.477 0 0 4.484 0 10.017c0 4.425 2.865 8.18 6.839 9.504.5.092.682-.217.682-.483 0-.237-.008-.868-.013-1.703-2.782.605-3.369-1.343-3.369-1.343-.454-1.158-1.11-1.466-1.11-1.466-.908-.62.069-.608.069-.608 1.003.07 1.531 1.032 1.531 1.032.892 1.53 2.341 1.088 2.91.832.092-.647.35-1.088.636-1.338-2.22-.253-4.555-1.113-4.555-4.951 0-1.093.39-1.988 1.029-2.688-.103-.253-.446-1.272.098-2.65 0 0 .84-.27 2.75 1.026A9.564 9.564 0 0110 4.844c.85.004 1.705.115 2.504.337 1.909-1.296 2.747-1.027 2.747-1.027.546 1.379.203 2.398.1 2.651.64.7 1.028 1.595 1.028 2.688 0 3.848-2.339 4.695-4.566 4.942.359.31.678.921.678 1.856 0 1.338-.012 2.419-.012 2.747 0 .268.18.58.688.482A10.019 10.019 0 0020 10.017C20 4.484 15.522 0 10 0z"
                                                            clipRule="evenodd"
                                                        />
                                                    </svg>
                                                </div>
                                            </div>
                                        </div>
                                    </div>

                                    <div className="mt-6 relative">
                                        <div className="absolute inset-0 flex items-center" aria-hidden="true">
                                            <div className="w-full border-t border-gray-300" />
                                        </div>
                                        <div className="relative flex justify-center text-sm">
                                            <span className="px-2 bg-white text-gray-500">Or</span>
                                        </div>
                                    </div>

                                    <div className="mt-6">
                                        <div className="space-y-6">
                                            <div>
                                                <label htmlFor="mobile-or-email" className="sr-only">
                                                    Email
                                                </label>
                                                <input
                                                    type="email"
                                                    name="email"
                                                    id="email"
                                                    autoComplete="email"
                                                    placeholder="Email"
                                                    value={email}
                                                    onChange={handleEmailChange}
                                                    required
                                                    className="block w-full shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm border-gray-300 rounded-md"
                                                />
                                            </div>

                                            <div>
                                                <label htmlFor="password" className="sr-only">
                                                    Password
                                                </label>
                                                <input
                                                    id="password"
                                                    name="password"
                                                    type="password"
                                                    placeholder="Password"
                                                    autoComplete="current-password"
                                                    value={password}
                                                    onChange={handlePasswordChange}
                                                    required
                                                    className="block w-full shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm border-gray-300 rounded-md"
                                                />
                                            </div>

                                            <div>
                                                <button
                                                    onClick={() => firebase.auth().signInWithEmailAndPassword(email, password).then(function(res){
                                                        history.replace("/");
                                                    }).catch(function(err){
                                                        setError(err.message)
                                                    })}
                                                    className="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
                                                >
                                                    Sign in
                                                </button>
                                            </div>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </main>
            </div>
        </div>
    )
}