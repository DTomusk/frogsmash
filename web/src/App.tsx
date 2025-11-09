import { Route, Routes } from "react-router-dom";
import Comparison from "./components/pages/Comparison";
import Template from "./components/templates/Template";
import LoadingPage from "./components/pages/LoadingPage";
import LeaderboardPage from "./components/pages/LeaderboardPage";

function App() {
  return (<Routes>
    <Route element={<Template />}>
        <Route path='/' element={<Comparison />} />
        <Route path='/loading' element={<LoadingPage />} />
        <Route path='/leaderboard' element={<LeaderboardPage />} />
    </Route>
  </Routes>
  );
}

export default App;
