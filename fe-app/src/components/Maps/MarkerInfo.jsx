import React, { useState } from 'react';
import Modal from 'react-bootstrap/Modal';
import axios from 'axios';

export default function MarkerInfo({ show, setShow }) {
  
  const handleClose = () => setShow(false);
  
  return (
    <>
      <Modal show={show} onHide={handleClose} animation={false} className="location-modal mx-auto" centered>
        <Modal.Header closeButton> 
          <Modal.Title>Marker data</Modal.Title>
        </Modal.Header>
                
        <Modal.Body>

        </Modal.Body>
      </Modal>
    </>
  )
}
