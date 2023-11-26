import React, { useState } from 'react';
import Modal from 'react-bootstrap/Modal';

export default function MarkerInfo({ show, setShow, markerPost}) {


  const [markerPosts, setMarkerPosts] = useState()
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
