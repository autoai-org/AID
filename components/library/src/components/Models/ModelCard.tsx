import { Fragment, useState } from 'react'
import { getInitials } from '../../services/utility'
import { Menu, Transition, Dialog } from '@headlessui/react'
import {
    InformationCircleIcon,
    ChatAltIcon,
    CodeIcon,
    DotsVerticalIcon,
    EyeIcon,
    FlagIcon,
    ShareIcon,
    DownloadIcon,
    StarIcon,
    ThumbUpIcon,
    QuestionMarkCircleIcon,
} from '@heroicons/react/solid'

const question =
{
    id: '81614',
    likes: '29',
    replies: '11',
    views: '2.7k',
    author: {
        name: 'Dries Vincent',
        imageUrl:
            'https://images.unsplash.com/photo-1506794778202-cad84cf45f1d?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80',
        href: '#',
    },
    date: 'December 9 at 11:43 AM',
    datetime: '2020-12-09T11:43:00',
    href: '#',
    title: 'What would you have done differently if you ran Jurassic Park?',
    body: `
      <p>Jurassic Park was an incredible idea and a magnificent feat of engineering, but poor protocols and a disregard for human safety killed what could have otherwise been one of the best businesses of our generation.</p>
      <p>Ultimately, I think that if you wanted to run the park successfully and keep visitors safe, the most important thing to prioritize would be&hellip;</p>
    `,
}
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
        console.log('set')
    }

    return (
        <ul className="py-4 space-y-2 sm:px-6 sm:space-y-4 lg:px-8">
            <li className="bg-white px-4 py-6 shadow sm:rounded-lg sm:px-6">
                <div>
                    <div className="flex space-x-3">
                        <div className="flex-shrink-0">
                            {!(model.avatar) ? (
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
                                                                href="#"
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
                                                                href="#"
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
                                                                href="#"
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
                                                                href="#"
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
                    {model.analysis.requirements.map((item: any, idx: number) => (
                        <span
                            className={classNames(
                                tagcolor[Math.ceil(idx % 5)],
                                'inline-flex items-center px-3 py-0.5 rounded-full text-sm font-medium mr-1 mt-8'
                            )}
                        >
                            {item}
                        </span>
                    ))}
                </div>
                <div className="mt-6 flex justify-between space-x-8">
                    <div className="flex space-x-6">
                        <span className="inline-flex items-center text-sm">
                            <button className="inline-flex space-x-2 text-gray-400 hover:text-gray-500">
                                <ThumbUpIcon className="h-5 w-5" aria-hidden="true" />
                                <span className="font-medium text-gray-900">{question.likes}</span>
                                <span className="sr-only">likes</span>
                            </button>
                        </span>
                        <span className="inline-flex items-center text-sm">
                            <button className="inline-flex space-x-2 text-gray-400 hover:text-gray-500">
                                <ChatAltIcon className="h-5 w-5" aria-hidden="true" />
                                <span className="font-medium text-gray-900">{question.replies}</span>
                                <span className="sr-only">replies</span>
                            </button>
                        </span>
                        <span className="inline-flex items-center text-sm">
                            <button className="inline-flex space-x-2 text-gray-400 hover:text-gray-500">
                                <EyeIcon className="h-5 w-5" aria-hidden="true" />
                                <span className="font-medium text-gray-900">{question.views}</span>
                                <span className="sr-only">views</span>
                            </button>
                        </span>
                    </div>
                    <div className="flex text-sm">
                        <span className="inline-flex items-center text-sm mr-2">
                            <a href={"/model/"+model.vendor+"/"+model.name} className="inline-flex space-x-2 text-gray-400 hover:text-gray-500">
                                <InformationCircleIcon className="h-5 w-5" aria-hidden="true" />
                                <span className="font-medium text-gray-900">Details</span>
                            </a>
                        </span>
                        <span className="inline-flex items-center text-sm">
                            <button onClick={handlePopup} className="inline-flex space-x-2 text-gray-400 hover:text-gray-500">
                                <DownloadIcon className="h-5 w-5" aria-hidden="true" />
                                <span className="font-medium text-gray-900">Install</span>
                            </button>
                        </span>

                        <Transition.Root show={open} as={Fragment}>
                            <Dialog as="div" static className="fixed z-10 inset-0 overflow-y-auto" open={open} onClose={setOpen}>
                                <div className="flex items-end justify-center min-h-screen pt-4 px-4 pb-20 text-center sm:block sm:p-0">
                                    <Transition.Child
                                        as={Fragment}
                                        enter="ease-out duration-300"
                                        enterFrom="opacity-0"
                                        enterTo="opacity-100"
                                        leave="ease-in duration-200"
                                        leaveFrom="opacity-100"
                                        leaveTo="opacity-0"
                                    >
                                        <Dialog.Overlay className="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity" />
                                    </Transition.Child>

                                    {/* This element is to trick the browser into centering the modal contents. */}
                                    <span className="hidden sm:inline-block sm:align-middle sm:h-screen" aria-hidden="true">
                                        &#8203;
                                    </span>
                                    <Transition.Child
                                        as={Fragment}
                                        enter="ease-out duration-300"
                                        enterFrom="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"
                                        enterTo="opacity-100 translate-y-0 sm:scale-100"
                                        leave="ease-in duration-200"
                                        leaveFrom="opacity-100 translate-y-0 sm:scale-100"
                                        leaveTo="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"
                                    >
                                        <div className="inline-block align-bottom bg-white rounded-lg px-4 pt-5 pb-4 text-left overflow-hidden shadow-xl transform transition-all sm:my-8 sm:align-middle sm:max-w-sm sm:w-full sm:p-6">
                                            <div>
                                                <div className="mx-auto flex items-center justify-center h-12 w-12 rounded-full bg-yellow-100">
                                                    <QuestionMarkCircleIcon className="h-6 w-6 text-yellow-600" aria-hidden="true" />
                                                </div>
                                                <div className="mt-3 text-center sm:mt-5">
                                                    <Dialog.Title as="h3" className="text-lg leading-6 font-medium text-gray-900">
                                                        {'Install ' + model.name}
                                                    </Dialog.Title>
                                                    <div className="mt-2">
                                                        <p className="text-sm text-gray-500">
                                                            {'Run the command below in your terminal:'}
                                                        </p>
                                                        <div className="my-1 text-xl font-medium text-gray-900">

                                                            {model.githubURL}

                                                        </div>
                                                    </div>
                                                </div>
                                            </div>
                                            <div className="mt-5 sm:mt-6">
                                                <button
                                                    type="button"
                                                    className="inline-flex justify-center w-full rounded-md border border-transparent shadow-sm px-4 py-2 bg-indigo-600 text-base font-medium text-white hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 sm:text-sm"
                                                    onClick={() => setOpen(false)}
                                                >
                                                    Close
                                                </button>
                                            </div>
                                        </div>
                                    </Transition.Child>
                                </div>
                            </Dialog>
                        </Transition.Root>


                    </div>
                </div>
            </li>
        </ul>
    )
}