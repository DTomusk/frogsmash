import { Route, Routes } from "react-router-dom";
import Comparison from "./components/Comparison";
import Template from "./components/Template";

function App() {
  return (<Routes>
    <Route element={<Template />}>
        <Route path='/' element={<Comparison />} />
    </Route>
  </Routes>
  );
}

export default App;
