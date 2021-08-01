// Copyright (c) 2021 Xiaozhe Yao et al.
// 
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

import MainLayout from "../Layouts/MainLayout";
import Profile from '../components/Profile/Profile'
export default function SearchResult() {

    return (
        <MainLayout>
            <div className="max-w-screen-xl mx-auto pb-6 px-4 sm:px-6 lg:pb-16 lg:px-8 pt-4">
                <Profile />
            </div>
        </MainLayout>
    )
}