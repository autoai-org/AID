// Copyright (c) 2021 Xiaozhe Yao et al.
// 
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

import { CheckIcon } from '@heroicons/react/solid'

const steps = [
    { name: 'Form Request', description: 'Compose the requests.', href: '#', status: 'current' },
    {
        name: 'View Results',
        description: 'View the results from the solver.',
        href: '#',
        status: 'upcoming',
    },
    { name: 'Detailed Analysis', description: 'View detailed analysis of the request.', href: '#', status: 'upcoming' },
]

function classNames(...classes: any) {
    return classes.filter(Boolean).join(' ')
}

export default function StepIndicator(props:any) {
    const currentStep = props.currentStep
    const currentStepIndex = steps.findIndex((s => s.name == currentStep))
    if(currentStepIndex != 0) {
        steps[currentStepIndex - 1].status = 'complete'
        steps[currentStepIndex].status = 'current'
    }
    
    return ( 
        <nav aria-label="Progress">
            <ol className="overflow-hidden">
                {steps.map((step, stepIdx) => (
                    <li key={step.name} className={classNames(stepIdx !== steps.length - 1 ? 'pb-10' : '', 'relative')}>
                        {step.status === 'complete' ? (
                            <>
                                {stepIdx !== steps.length - 1 ? (
                                    <div className="-ml-px absolute mt-0.5 top-4 left-4 w-0.5 h-full bg-indigo-600" aria-hidden="true" />
                                ) : null}
                                <a href={step.href} className="relative flex items-start group">
                                    <span className="h-9 flex items-center">
                                        <span className="relative z-10 w-8 h-8 flex items-center justify-center bg-indigo-600 rounded-full group-hover:bg-indigo-800">
                                            <CheckIcon className="w-5 h-5 text-white" aria-hidden="true" />
                                        </span>
                                    </span>
                                    <span className="ml-4 min-w-0 flex flex-col">
                                        <span className="text-xs font-semibold tracking-wide uppercase">{step.name}</span>
                                        <span className="text-sm text-gray-500">{step.description}</span>
                                    </span>
                                </a>
                            </>
                        ) : step.status === 'current' ? (
                            <>
                                {stepIdx !== steps.length - 1 ? (
                                    <div className="-ml-px absolute mt-0.5 top-4 left-4 w-0.5 h-full bg-gray-300" aria-hidden="true" />
                                ) : null}
                                <a href={step.href} className="relative flex items-start group" aria-current="step">
                                    <span className="h-9 flex items-center" aria-hidden="true">
                                        <span className="relative z-10 w-8 h-8 flex items-center justify-center bg-white border-2 border-indigo-600 rounded-full">
                                            <span className="h-2.5 w-2.5 bg-indigo-600 rounded-full" />
                                        </span>
                                    </span>
                                    <span className="ml-4 min-w-0 flex flex-col">
                                        <span className="text-xs font-semibold tracking-wide uppercase text-indigo-600">{step.name}</span>
                                        <span className="text-sm text-gray-500">{step.description}</span>
                                    </span>
                                </a>
                            </>
                        ) : (
                            <>
                                {stepIdx !== steps.length - 1 ? (
                                    <div className="-ml-px absolute mt-0.5 top-4 left-4 w-0.5 h-full bg-gray-300" aria-hidden="true" />
                                ) : null}
                                <a href={step.href} className="relative flex items-start group">
                                    <span className="h-9 flex items-center" aria-hidden="true">
                                        <span className="relative z-10 w-8 h-8 flex items-center justify-center bg-white border-2 border-gray-300 rounded-full group-hover:border-gray-400">
                                            <span className="h-2.5 w-2.5 bg-transparent rounded-full group-hover:bg-gray-300" />
                                        </span>
                                    </span>
                                    <span className="ml-4 min-w-0 flex flex-col">
                                        <span className="text-xs font-semibold tracking-wide uppercase text-gray-500">{step.name}</span>
                                        <span className="text-sm text-gray-500">{step.description}</span>
                                    </span>
                                </a>
                            </>
                        )}
                    </li>
                ))}
            </ol>
        </nav>
    )
}