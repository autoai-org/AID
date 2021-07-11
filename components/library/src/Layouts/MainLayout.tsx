// Copyright (c) 2021 Xiaozhe Yao et al.
// 
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

import AIDHeader from "../components/Header";
import AIDFooter from "../components/Footer";

export default function MainLayout(props: any) {
    return (
        <div className="bg-white">
            <div className="relative overflow-hidden">
                <AIDHeader />
                {props.children}
                <AIDFooter />
            </div>
        </div>
    );
}