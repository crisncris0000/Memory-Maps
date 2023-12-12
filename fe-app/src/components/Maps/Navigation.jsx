import React from 'react';
import Container from 'react-bootstrap/Container';
import Navbar from 'react-bootstrap/Navbar';
import Nav from 'react-bootstrap/Nav';
import NavDropdown from 'react-bootstrap/NavDropdown';
import { Link } from "react-router-dom";
import "../../css/navbar.css";

export default function Navigation() {
    return (
        <Navbar className="bg-body-tertiary custom-navbar" expand="lg">
            <Container>
                <Link className="nav-link" to="/"><Navbar.Brand>Nostalgia Maps</Navbar.Brand></Link>
                <Navbar.Toggle />
                <Navbar.Collapse className="justify-content-end">
                    <NavDropdown title="Friends List" id="basic-nav-dropdown" className="text-center">
                        <NavDropdown.Item href="#action/3.1">Action</NavDropdown.Item>
                        <NavDropdown.Item href="#action/3.2">Another action</NavDropdown.Item>
                        <NavDropdown.Item href="#action/3.3">Something</NavDropdown.Item>
                        <NavDropdown.Divider />
                        <NavDropdown.Item href="#action/3.4">Separated link</NavDropdown.Item>
                    </NavDropdown>
                    <Nav.Link className='logout'>Logout</Nav.Link>
                </Navbar.Collapse>
                
            </Container>
        </Navbar>
    );
}
