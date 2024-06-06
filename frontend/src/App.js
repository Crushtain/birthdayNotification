import { Container, Row, Col  } from "react-bootstrap"
import React, {useEffect, useState} from "react";
import axios from "axios";
import "./App.css"
import {Route ,Routes} from "react-router-dom";
import Home from "./page/Home";
import SingleUser from "./page/SingleUser";
import Header from "./components/Header";

function App() {



    return (
        <>
            <Header/>
            <Routes>
                <Route path="/" element={<Home />}/>
                <Route path="/user/:id" element={<SingleUser/>}></Route>
            </Routes>

        </>

  );
}

export default App;
