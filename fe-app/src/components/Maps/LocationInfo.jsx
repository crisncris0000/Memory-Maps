import React, { useState } from 'react';
import Button from 'react-bootstrap/Button';
import Modal from 'react-bootstrap/Modal';

export default function LocationInfo({show, setShow}) {
    
    const handleClose = () => setShow(false);

    return (
        <>
            <Modal show={show}
                onHide={handleClose}
                animation={false}>
                <Modal.Header closeButton>
                    <Modal.Title>Marker Form</Modal.Title>
                </Modal.Header>
                <Modal.Body>
                    <form>
                        <label>Image upload * </label>
                        <input type="file" required />
                        
                        <label>Description *</label>
                        <textarea type="text" placeholder="Please enter description" className="description"/>
                    </form>
                </Modal.Body>
                <Modal.Footer>
                    <Button variant="secondary"
                        onClick={handleClose}>
                        Close
                    </Button>
                    <Button variant="primary"
                        onClick={handleClose}>
                        Submit
                    </Button>
                </Modal.Footer>
            </Modal>
        </>
    )
}
