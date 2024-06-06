import React, {useEffect, useState} from "react";
import axios from "axios";
import {Link, useParams} from "react-router-dom";
import {Col, Container, Row} from "react-bootstrap";


const SingleUser = () => {

    const params = useParams()
    const [apiData, setApiData] = useState(false)
    useEffect(() => {
        const fetchData = async () => {
            try {
                const apiUrl = process.env.REACT_APP_API_ROOT+ "/" + params.id;
                const response = await axios.get(apiUrl);

                if (response.status === 201) {
                    setApiData(response?.data?.user);
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
       <Container>
           <Row>
               {/*{apiData &&*/}
               {/*    apiData.map((user, index) => (*/}
               {/*        <Col key={index} xs="4" className="py-5 box">*/}
               {/*            <div className="box-title">{user.username}</div>*/}
               {/*            <div>{user.birthday}</div>*/}
               {/*        </Col>*/}
               {/*    ))}*/}
               <Col xs="12">{apiData.username}</Col>
               <Col xs="12">{apiData.birthday}</Col>
           </Row>
       </Container>
    )
}


export default SingleUser