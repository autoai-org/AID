import MainLayout from "../Layouts/MainLayout"
import PublishForm from '../components/Models/PublishForm'
export default function Publish() {
    return (
        <MainLayout>
            <div className="max-w-screen-xl mx-auto pb-6 px-4 sm:px-6 lg:pb-16 lg:px-8 pt-4">
                <PublishForm/>
            </div>
        </MainLayout>
    )
}