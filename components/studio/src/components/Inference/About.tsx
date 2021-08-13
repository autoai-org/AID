// Copyright (c) 2021 Xiaozhe Yao et al.
// 
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

import Divider from '../Common/Divider'

export default function AboutRepository(props:any) {
    console.log(props)
    return (
        <div className="md:grid md:grid-cols-3 md:gap-6">
          <div className="md:col-span-2">
            <div className="px-4 sm:px-0">
            <h2 className="text-lg font-medium leading-6 text-gray-900">Request Info</h2>
              You are about to send request to <span className="text-black font-bold">{props.solver.name}</span> solver.
              <Divider/>
              <h3 className="text-lg font-medium leading-6 text-gray-900">Privacy</h3>
              <p>
                Your request, including the content of the form and the statistics data, will be uploaded to the server your connected to. Please be sure that you trust the server (and the network between you and the server, if you are not on an encrypted network).
                
                Read our <a href="https://aid.autoai.org/docs/about/privacy" className="text-gray-400 hover:text-gray-500">privacy policy</a>.
              </p>
            </div>
          </div>
          </div>
    )
}