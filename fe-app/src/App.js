import './css/App.css';
import Home from './components/Home/Home';
import 'bootstrap/dist/css/bootstrap.min.css';
import 'mapbox-gl/dist/mapbox-gl.css';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Login from './components/Authentication/Login';
import Maps from './components/Maps/Maps';
import Register from './components/Authentication/Register';

function App() {
  return (
    <Router>
      <div className="App">
      <Routes>
        <Route path='/' element={<Home />}></Route>
        <Route path='/login' element={<Login />}></Route>
        <Route path='/maps' element={<Maps />}></Route>
        <Route path='/register' element={<Register />}></Route>
      </Routes>
      </div>
    </Router>
  );
}

export default App;