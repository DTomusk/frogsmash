import { Route, Routes } from "react-router-dom";
import Comparison from "./components/pages/Comparison";
import Template from "./components/templates/Template";
import LoadingPage from "./components/pages/LoadingPage";
import LeaderboardPage from "./components/pages/LeaderboardPage";
import NotFoundPage from "./components/pages/NotFoundPage";
import UploadPage from "./components/pages/UploadPage";

function App() {
  return (<Routes>
    <Route element={<Template />}>
        <Route path='/' element={<Comparison />} />
        <Route path='/loading' element={<LoadingPage />} />
        <Route path='/leaderboard' element={<LeaderboardPage />} />
        <Route path='/upload' element={<UploadPage />} />
        <Route path='*' element={<NotFoundPage />} />
    </Route>
  </Routes>
  );
}

export default App;
