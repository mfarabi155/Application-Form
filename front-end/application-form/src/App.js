import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import 'bootstrap/dist/css/bootstrap.min.css';
import PersonTable from './PersonTable';
import PersonForm from './PersonForm';

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<PersonTable />} />
        <Route path="/add-person" element={<PersonForm />} />
        <Route path="/edit-person/:identityNumber" element={<PersonForm />} />
      </Routes>
    </Router>
  );
}

export default App;
