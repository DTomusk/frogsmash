import { Route, Routes } from "react-router-dom";
import Comparison from "./components/Comparison";
import Template from "./components/Template";
import LoadingSpinner from "./components/LoadingSpinner";

function App() {
  return (<Routes>
    <Route element={<Template />}>
        <Route path='/' element={<Comparison />} />
        <Route path='/loading' element={<LoadingSpinner />} />
    </Route>
  </Routes>
  );
}

export default App;
