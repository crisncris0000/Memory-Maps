import React, { useEffect, useState } from 'react';
import Modal from 'react-bootstrap/Modal';
import axios from 'axios';
import Carousel from 'react-bootstrap/Carousel';
import { useSelector } from 'react-redux';


export default function MarkerInfo({ show, setShow, markerPost }) {
  const [showComments, setShowComments] = useState(false);

  return (
    <>
      {showComments === false ? 
      <MarkerPost show={show} setShow={setShow} markerPost={markerPost} showComments={showComments} setShowComments={setShowComments}/> 
      : 
      <MarkerComments show={show} setShow={setShow} markerPost={markerPost} setShowComments={setShowComments}/>}
    </>
  );
}

function MarkerPost({ show, setShow, markerPost, setShowComments }) {

  const handleClose = () => setShow(false);
  const [likes, setLikes] = useState(0);

  const [imagesInfo, setImagesInfo] = useState([]);
  const [postCreatedBy, setPostCreatedBy] = useState(null);

  const user = useSelector((state) => state.user.value);

  useEffect(() => {
    if (show) {
      axios.get(`http://localhost:8080/marker-post/images/${markerPost.id}`)
        .then((response) => {
          setImagesInfo(response.data.images);
        })
        .catch((error) => {
          console.log(error);
        });
    }
  }, [show]);

  const handleLikesClick = () => {
    setLikes(likes + 1);

    if(likes % 5 === 0) {
      axios.put(`http://localhost:8080/marker-posts/update`, {
        id: markerPost.id,
        likes
      }).then((response) => {
        console.log(response.data);
      }).catch((error) => {
        console.log(error);
      })
    }
  }

  const handleDeleteClick = () => {
    axios.delete(`http://localhost:8080/marker-posts/delete/${markerPost.id}`)
    .then((response) => {
      console.log(response.data);
    }).catch((error) => {
      console.log(error);
    })
  }

  return (
    <>
      {show && (
        <Modal show={show} onHide={handleClose} animation={false} className="location-modal mx-auto" centered>
          <Modal.Header closeButton>
            <Modal.Title>
              Post created by {postCreatedBy ? postCreatedBy.firstName + " " + postCreatedBy.lastName : null}
            </Modal.Title>
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
              <button className="likes" style={{ backgroundColor: "transparent" }} onClick={handleLikesClick}>
                <span role="img" aria-label="likes">👍</span> {markerPost.likes} Likes
              </button>
              <button className="comments" style={{ backgroundColor: "transparent" }} onClick={() => setShowComments(true)}>
                <span role="img" aria-label="comments">💬</span> {markerPost.comments} Comments
              </button>

              <button className="delete" style={{ backgroundColor: "transparent", color: "red" }} onClick={handleDeleteClick}>
                {markerPost.userID === user.id ? <span role="img" aria-label="delete">🗑️ Delete</span> : null}
              </button>
            </div>
          </Modal.Footer>
        </Modal>
      )}
    </>
  );
}  

function MarkerComments({ show, markerPost, setShow, setShowComments }) {

  const [comments, setComments] = useState([]);
  const [newComment, setNewComment] = useState('');

  const user = useSelector((state) => state.user.value);

  const handleClose = () => {
    setShow(false);
    setShowComments(false);
  }

  const handleAddComment = () => {
    axios.post("http://localhost:8080/comments/new", {
      userID: user.id,
      markerID: markerPost.id,
      comment: newComment,
      likes: 0,
    }).then((response) => {
      console.log(response.data);
      setComments((prevComments) => 
        [...prevComments, 
          {
            firstName: user.firstName, 
            lastName: user.lastName, 
            comment: response.data.comment.comment,
            email: user.email
          }
        ]);
    }).catch((error) => {
      console.log(error);
    })
  }

  const handleReturn = () => {
    setShowComments(false);
  }

  const handleDelete = (id) => {
    axios.delete(`http://localhost:8080/comments/delete/${id}`)
    .then((response) => {
      console.log(response.data);
      const filteredArr = comments.filter((comment) => (
        comment.id !== id
      ));
      setComments(filteredArr);
    }).catch((error) => {
      console.log(error);
    })
  }

  useEffect(() => {
    axios.get(`http://localhost:8080/comments/${markerPost.id}`)
      .then((response) => {
        setComments(response.data.comments);
      })
      .catch((error) => {
        console.log(error);
      });
  }, []);

  return (
    <>
      {show &&
        <Modal show={show} onHide={handleClose} animation={false} className="location-modal mx-auto" centered>
          <Modal.Header closeButton>
            <Modal.Title className="text-center">Marker Comments</Modal.Title>
          </Modal.Header>

          <Modal.Body>
            <div className="comments-container">
              {comments ? comments.map((userComment) => (
                <div className="comment" key={userComment.id}>
                  <div className="comment-header">
                    <strong>{`${userComment.firstName} ${userComment.lastName}`}</strong>
                    {user.email === userComment.email ? 
                    <button className="btn btn-sm btn-danger delete-button" onClick={() => handleDelete(userComment.id)}>Delete</button> 
                    : null}
                  </div>
                  <div className="comment-text">{`${userComment.comment}`}</div>
                </div>              
              )) : null}

            </div>
          </Modal.Body>



          <Modal.Footer>
            <div className="footer-content">
              <button className="btn btn-primary" onClick={handleReturn}>Return</button>
              <div className="comment-input">
                <input
                  type="text"
                  placeholder="Add a comment..."
                  className="form-control"
                  onChange={(e) => setNewComment(e.target.value)}
                />
              </div>
              <div className="comment-button">
                <button onClick={handleAddComment} className="btn btn-primary">
                  Send
                </button>
              </div>
            </div>
          </Modal.Footer>
        </Modal>
      }
    </>
  );
}

