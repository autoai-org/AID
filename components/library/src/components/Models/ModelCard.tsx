import { Fragment, useState } from 'react'
import { getInitials } from '../../services/utility'
import { Menu, Transition } from '@headlessui/react'
import {
    InformationCircleIcon,
    CodeIcon,
    DotsVerticalIcon,
    ArrowDownIcon,
    FlagIcon,
    ShareIcon,
    DownloadIcon,
    StarIcon,
    ThumbUpIcon,
} from '@heroicons/react/solid'
import InstallModelDialog from './InstallModelDialog'
const tagcolor = [
    'bg-indigo-100 text-indigo-800', 'bg-pink-100 text-pink-800',
    'bg-green-100 text-green-800', 'bg-yellow-100 text-yellow-800',
    'bg-blue-100 text-blue-800'
]
function classNames(...classes: any) {
    return classes.filter(Boolean).join(' ')
}
export default function ModelCard(props: any) {
    const model = props.model
    const [open, setOpen] = useState(false)

    const handlePopup = () => {
        setOpen(true)
    }
    function handleClose() {
        setOpen(false)
    }
    return (

        <ul className="py-4 space-y-2 sm:px-6 sm:space-y-4 lg:px-8">
            <li className="bg-white px-4 py-6 shadow sm:rounded-lg sm:px-6">
                <div>
                    <div className="flex space-x-3">
                        <div className="flex-shrink-0">
                            {(model.avatar || !(model.avatar ==="null") || !(model.avatar ==="undefined")) ? (
                                <img className="h-10 w-10 rounded-full" src={model.avatar} alt="" />
                            ) :
                                (
                                    <span className="inline-flex items-center justify-center h-10 w-10 rounded-full bg-gray-500">
                                        <span className="font-medium leading-none text-white">{getInitials(model.name)}</span>
                                    </span>
                                )}

                        </div>
                        <div className="min-w-0 flex-1">
                            <div className="my-1 text-2xl font-medium text-gray-900">

                                {model.vendor + "/" + model.name}

                            </div>
                        </div>
                        <div className="flex-shrink-0 self-center flex">
                            <Menu as="div" className="relative inline-block text-left">
                                {({ open }) => (
                                    <>
                                        <div>
                                            <Menu.Button className="-m-2 p-2 rounded-full flex items-center text-gray-400 hover:text-gray-600">
                                                <span className="sr-only">Open options</span>
                                                <DotsVerticalIcon className="h-5 w-5" aria-hidden="true" />
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
                                                className="origin-top-right absolute right-0 mt-2 w-56 rounded-md shadow-lg bg-white ring-1 ring-black ring-opacity-5 focus:outline-none"
                                            >
                                                <div className="py-1">
                                                    <Menu.Item>
                                                        {({ active }) => (
                                                            <a
                                                                href="/"
                                                                className={classNames(
                                                                    active ? 'bg-gray-100 text-gray-900' : 'text-gray-700',
                                                                    'flex px-4 py-2 text-sm'
                                                                )}
                                                            >
                                                                <StarIcon className="mr-3 h-5 w-5 text-gray-400" aria-hidden="true" />
                                                                <span>Add to favorites</span>
                                                            </a>
                                                        )}
                                                    </Menu.Item>
                                                    <Menu.Item>
                                                        {({ active }) => (
                                                            <a
                                                                href="/"
                                                                className={classNames(
                                                                    active ? 'bg-gray-100 text-gray-900' : 'text-gray-700',
                                                                    'flex px-4 py-2 text-sm'
                                                                )}
                                                            >
                                                                <CodeIcon className="mr-3 h-5 w-5 text-gray-400" aria-hidden="true" />
                                                                <span>Embed</span>
                                                            </a>
                                                        )}
                                                    </Menu.Item>
                                                    <Menu.Item>
                                                        {({ active }) => (
                                                            <a
                                                                href="/"
                                                                className={classNames(
                                                                    active ? 'bg-gray-100 text-gray-900' : 'text-gray-700',
                                                                    'flex px-4 py-2 text-sm'
                                                                )}
                                                            >
                                                                <FlagIcon className="mr-3 h-5 w-5 text-gray-400" aria-hidden="true" />
                                                                <span>Report content</span>
                                                            </a>
                                                        )}
                                                    </Menu.Item>
                                                    <Menu.Item>
                                                        {({ active }) => (
                                                            <a
                                                                href="/"
                                                                className={classNames(
                                                                    active ? 'bg-gray-100 text-gray-900' : 'text-gray-700',
                                                                    'flex px-4 py-2 text-sm'
                                                                )}
                                                            >
                                                                <ShareIcon className="mr-3 h-5 w-5 text-gray-400" aria-hidden="true" />
                                                                <span>Share</span>
                                                            </a>
                                                        )}
                                                    </Menu.Item>
                                                </div>
                                            </Menu.Items>
                                        </Transition>
                                    </>
                                )}
                            </Menu>
                        </div>
                    </div>
                    <h2 id={'models-description' + model._id} className="mt-4 text-base font-medium text-gray-900">
                        {model.description}
                    </h2>
                </div>
                <div>
                    {model.analysis.requirements.map((item: any, idx: number) => {
                        item = item.replace(/(\r\n|\n|\r)/gm, "");
                        if (!(["mlpm","cython","h5py"].includes(item))) {
                            return(
                                <span
                                    key={'requirements-' + idx}
                                    className={classNames(
                                        tagcolor[Math.ceil(idx % 5)],
                                        'inline-flex items-center px-3 py-0.5 rounded-full text-sm font-medium mr-1 mt-8'
                                    )}
                                >
                                    {item}
                                </span>
                            )
                        } else {return null}
                    })}
                </div>
                <div className="mt-6 flex justify-between space-x-8">
                    <div className="flex space-x-6">
                        <span className="inline-flex items-center text-sm">
                            <button className="inline-flex space-x-2 text-gray-400 hover:text-gray-500">
                                <ThumbUpIcon className="h-5 w-5" aria-hidden="true" />
                                <span className="font-medium text-gray-900">{model.analysis.summary.stargazers_count + model.analysis.summary.subscribers_count + model.analysis.summary.watchers_count}</span>
                                <span className="sr-only">likes</span>
                            </button>
                        </span>
                        <span className="inline-flex items-center text-sm">
                            <button className="inline-flex space-x-2 text-gray-400 hover:text-gray-500">
                                <ArrowDownIcon className="h-5 w-5" aria-hidden="true" />
                                <div className='has-tooltip'>
                                <span className='tooltip rounded shadow-lg p-1 bg-gray-100 text-gray-500 -mt-8'>We have not collected the installation statistics yet, stay tuned.</span>
                                N/A
                                </div>
                                <span className="sr-only">N/A</span>
                            </button>
                        </span>
                    </div>
                    <div className="flex text-sm">
                        <span className="inline-flex items-center text-sm mr-2">
                            <a href={"/model/" + model.vendor + "/" + model.name} className="inline-flex space-x-2 text-gray-400 hover:text-gray-500">
                                <InformationCircleIcon className="h-5 w-5" aria-hidden="true" />
                                <span className="font-medium text-gray-900">Details</span>
                            </a>
                        </span>
                        <span className="inline-flex items-center text-sm">
                            <button onClick={handlePopup} className="inline-flex space-x-2 text-gray-400 hover:text-gray-500">
                                <DownloadIcon className="h-5 w-5" aria-hidden="true" />
                                <span className="font-medium text-gray-900">Install</span>
                            </button>
                        </span>

                        <InstallModelDialog model={model} open={open} onClose={handleClose} />
                    </div>
                </div>
            </li>
        </ul>

    )
}