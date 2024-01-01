import React, { useEffect, useState } from 'react';
import Modal from 'react-bootstrap/Modal';
import axios from 'axios';
import Carousel from 'react-bootstrap/Carousel';


export default function MarkerInfo({ show, setShow, markerPost }) {
  const [showComments, setShowComments] = useState(false);

  return (
    <>
      {showComments === false ? 
      <MarkerPost show={show} setShow={setShow} markerPost={markerPost} showComments={showComments} setShowComments={setShowComments}/> 
      : 
      <MarkerComments show={show} setShow={setShow} showComments={showComments} setShowComments={setShowComments}/>}
    </>
  );
}

function MarkerPost({ show, setShow, markerPost, setShowComments }) {

  const handleClose = () => setShow(false);

  const [imagesInfo, setImagesInfo] = useState([]);

  useEffect(() => {
    if (show) {
      axios.get(`http://localhost:8080/marker-post/images/${markerPost.id}`)
        .then((response) => {
          setImagesInfo(response.data.images);
          console.log(markerPost);
        })
        .catch((error) => {
          console.log(error);
        });
    }
  }, [show]);

  return (
    <>
      {show &&
        <Modal show={show} onHide={handleClose} animation={false} className="location-modal mx-auto" centered>
          <Modal.Header closeButton>
            <Modal.Title>Marker Details</Modal.Title>
          </Modal.Header>
          <Modal.Body>
            <Carousel>
              {imagesInfo.map((image, index) => (
                <Carousel.Item key={index}>
                  <img
                    className="carousel-img"
                    src={`data:${image.mimeType};base64,${image.image}`}
                    alt={`Image ${index + 1}`}
                  />
                </Carousel.Item>
              ))}
            </Carousel>

            <div className="description">
              <h2>Description</h2>
              <p>{markerPost.description}</p>
            </div>
          </Modal.Body>

          <Modal.Footer>
            <div className="likes-comments">
              <button className="likes" style={{backgroundColor: "transparent"}}>
                <span role="img" aria-label="likes">👍</span> {markerPost.likes} Likes
              </button>
              <button className="comments" style={{backgroundColor: "transparent"}} onClick={() => setShowComments(true)}>
                <span role="img" aria-label="comments">💬</span> {markerPost.comments} Comments
              </button>
            </div>
          </Modal.Footer>
        </Modal>
      }
    </>
  );
}

function MarkerComments ({ show, setShow, showComments, setShowComments }) {

  const handleClose = () => setShow(false);

  return (
    <>
      {show &&
        <Modal show={show} onHide={handleClose} animation={false} className="location-modal mx-auto" centered>
          <Modal.Header closeButton>
            <Modal.Title>Marker Comments</Modal.Title>
          </Modal.Header>
          <Modal.Body>
          </Modal.Body>

          <Modal.Footer>
          </Modal.Footer>
        </Modal>
      }
    </>
  );
}
