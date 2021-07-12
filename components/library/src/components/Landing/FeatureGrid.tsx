import {
    CloudUploadIcon,
    CogIcon,
    LockClosedIcon,
    RefreshIcon,
    PaperAirplaneIcon,
    AcademicCapIcon
} from '@heroicons/react/outline'

const features = [
    {
        name: 'Upload Models',
        description: 'You can upload models to our database so that it is indexed, and others can quickly reproduce and rerun your models.',
        icon: CloudUploadIcon,
    },
    {
        name: 'Everything in Virtual Environment',
        description: 'Quickly construct virtual environment and put machine learning packages inside, so it has no harm to your server.',
        icon: LockClosedIcon,
    },
    {
        name: 'Quickly Updates',
        description: 'You have all the code to run the machine learning algorithms, which enables quickly edit it and re-build the environment.',
        icon: RefreshIcon,
    },
    {
        name: 'Extensible',
        description: 'Easily integrate machine learning in your project, and enable extensions to inspect or analyse machine learning models.',
        icon: PaperAirplaneIcon,
    },
    {
        name: 'Powerful API',
        description:
            'Access installed models, containers, solvers or packages via GraphQL or command line utility with ease.',
        icon: CogIcon,
    },
    {
        name: 'Open Science',
        description: 'Improve the reproducibility and transparentness by allowing others access your models and algorithms.',
        icon: AcademicCapIcon,
    },
]
export default function FeatureScreenShots() {
    return (
        <div className="relative bg-white py-16 sm:py-24 lg:py-32">
            <div className="mx-auto max-w-md px-4 text-center sm:max-w-3xl sm:px-6 lg:px-8 lg:max-w-7xl">
                <h2 className="text-base font-semibold tracking-wider text-cyan-600 uppercase">Access Machine Learning Faster</h2>
                <p className="mt-2 text-3xl font-extrabold text-gray-900 tracking-tight sm:text-4xl">
                    Everything you need to access machine learning
                </p>
                <p className="mt-5 max-w-prose mx-auto text-xl text-gray-500">
                    Easily download and deploy machine learning models, and then access it from command line, http requests, or web user interface.
                </p>
                <div className="mt-12">
                    <div className="grid grid-cols-1 gap-8 sm:grid-cols-2 lg:grid-cols-3">
                        {features.map((feature) => (
                            <div key={feature.name} className="pt-6">
                                <div className="flow-root bg-gray-50 rounded-lg px-6 pb-8">
                                    <div className="-mt-6">
                                        <div>
                                            <span className="inline-flex items-center justify-center p-3 bg-gradient-to-r from-teal-500 to-cyan-600 rounded-md shadow-lg">
                                                <feature.icon className="h-6 w-6 text-green" aria-hidden="true" />
                                            </span>
                                        </div>
                                        <h3 className="mt-8 text-lg font-medium text-gray-900 tracking-tight">{feature.name}</h3>
                                        <p className="mt-5 text-base text-gray-500">{feature.description}</p>
                                    </div>
                                </div>
                            </div>
                        ))}
                    </div>
                </div>
            </div>
        </div>
    )
}