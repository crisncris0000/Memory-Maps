import React from 'react';
import Button from 'react-bootstrap/Button';
import Modal from 'react-bootstrap/Modal';

export default function LocationInfo({ show, setShow, longitude, latitude, onHide }) {

    const handleClose = () => {
        onHide()
        setShow(false);

    }

    return (
        <>
            <Modal show={show} onHide={handleClose} animation={false} className="location-modal mx-auto" centered>
                <Modal.Header closeButton> 
                    <Modal.Title>Marker Form</Modal.Title>
                </Modal.Header>
                <Modal.Body>
                    <form>
                        <div className="form-group">
                            <label className="upload-label">Image upload *</label>
                            <input type="file" required className="form-control upload-input" />
                        </div>

                        <div className="form-group">
                            <label className="description-label">Description *</label>
                            <textarea type="text" placeholder="Please enter description" className="form-control description" />
                        </div>
                    </form>
                </Modal.Body>
                <Modal.Footer>
                    <Button variant="secondary" onClick={handleClose}>
                        Close
                    </Button>
                    <Button variant="primary" onClick={handleClose}>
                        Submit
                    </Button>
                </Modal.Footer>
            </Modal>
        </>
    );
}
