import { XCircleIcon } from '@heroicons/react/solid'

export default function Alert(props: any) {
    return (
        <div className="rounded-md bg-red-50 p-4" style={{ maxWidth: '80vh', marginLeft: 'auto', marginRight: 'auto', marginBottom:'60px' }}>
            <div className="flex">
                <div className="flex-shrink-0">
                    <XCircleIcon className="h-5 w-5 text-red-400" aria-hidden="true" />
                </div>
                <div className="ml-3">
                    <h3 className="text-sm font-medium text-red-800">{props.title}</h3>
                    <div className="mt-2 text-sm text-red-700">
                        {props.message}
                    </div>
                </div>
            </div>
        </div>
    )
}