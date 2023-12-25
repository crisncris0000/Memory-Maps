import './css/App.css';
import Home from './components/Home/Home';
import 'bootstrap/dist/css/bootstrap.min.css';
import 'mapbox-gl/dist/mapbox-gl.css';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Login from './components/Authentication/Login';
import Maps from './components/Maps/Maps';
import Register from './components/Authentication/Register';
import { Provider } from 'react-redux';
import store from './state/store';
import ResetPassword from './components/Authentication/ResetPassword';

function App() {
  return (
    <Router>
      <Provider store={store}>
        <div className="App">
          <Routes>
            <Route path='/' element={<Home />}></Route>
            <Route path='/login' element={<Login />}></Route>
            <Route path='/maps' element={<Maps />}></Route>
            <Route path='/register' element={<Register />}></Route>
            <Route path='/reset-password' element={<ResetPassword />}></Route>
          </Routes>
        </div>
      </Provider>
    </Router>
  );
}

export default App;