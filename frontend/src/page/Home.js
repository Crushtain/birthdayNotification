import React, {useEffect, useState} from "react";
import {Col, Container, Row} from "react-bootstrap";
import axios from "axios";
import {Last} from "react-bootstrap/PageItem";
import {Link} from "react-router-dom";


const Home = () => {

    const [apiData, setApiData] = useState(false);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const apiUrl = process.env.REACT_APP_API_ROOT;
                const response = await axios.get(apiUrl);

                if (response.status === 201) {
                    setApiData(response?.data?.users);
                }
            } catch (error){
                console.log(error.response);
            }



        };
        fetchData();


        return () => {};
    }, []);

    console.log(apiData)

    return (
        <Container className="py-2">
            <Row>
                {apiData &&
                    apiData.map((users, index) => (
                        <Col key={index} xs="4" className="py-5 box">
                            <div className="box-title">
                                <Link to={`user/${users.id}`}>{users.username}</Link>

                                </div>
                            <div>{users.birthday}</div>
                        </Col>
                    ))}
            </Row>
        </Container>
    )

}

export default Home