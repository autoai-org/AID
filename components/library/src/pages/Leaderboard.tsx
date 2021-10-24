import MainLayout from "../Layouts/MainLayout";
import Leaderboard from "../components/Profile/Leaderboard";

export default function LeaderboardPage() {
    return (
        <MainLayout>
            <div className="max-w-screen-xl mx-auto pb-6 px-4 sm:px-6 lg:pb-16 lg:px-8 pt-4">
                <Leaderboard/>
            </div>
        </MainLayout>
    )
}