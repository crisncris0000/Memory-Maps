import React from 'react';

export default function ContactForm() {
    return (
        <div className="contact-me">
            <h3>Technical issues regarding this website? Email me using this contact form:</h3>
            <form className="form-container">
                <div className="input-group">
                    <label htmlFor="name">Name</label>
                    <input type="text" id="name" name="name" required />
                </div>
                
                <div className="input-group">
                    <label htmlFor="email">Email</label>
                    <input type="email" id="email" name="email" required />
                </div>

                <div className="input-group">
                    <label htmlFor="message">Message</label>
                    <textarea id="message" name="message" rows="4" required></textarea>
                </div>
                
                <button type="submit">Send</button>
            </form>
        </div>
    );
}
