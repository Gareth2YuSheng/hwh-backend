import { SyntheticEvent, useEffect, useState } from "react";
import { useNavigate, NavLink } from "react-router-dom";
//Redux
import { useDispatch, useSelector } from 'react-redux';
import { fetchUserData } from "../store/userSlice";
import { AppDispatch, RootState } from '../store/store';
import { fetchThreadData } from "../store/threadSlice";
import { fetchTagData } from "../store/tagSlice";

import { Button, Pagination, Alert, Form } from "react-bootstrap";
import ArrowBackIcon from "@mui/icons-material/ArrowBackSharp";
import SearchIcon from '@mui/icons-material/Search';
import TagsDropdown from "../components/TagsDropdown";
import DisplayThreads from "../components/DisplayThreads";

import Cookies from "js-cookie";

export default function Home() {
  //Redux
  const { user } = useSelector((state: RootState) => state.user);
  const { error, totalThreads } = useSelector((state: RootState) => state.thread);
  const dispatch = useDispatch<AppDispatch>();

  const navigate = useNavigate();

  const [page, setPage] = useState(1);
  const [count, setCount] = useState(5);
  const [paginationItems, setPaginationItems] = useState<JSX.Element[]>([]);
  const [alertVisible, setAlertVisible] = useState(false);
  const [alertMessage, setAlertMessage] = useState<string | null>("");
  const [search, setSearch] = useState("");
  const [searching, setSearching] = useState(false);
  //For Tag Dropdown
  const [filterByTagID, setFilteredTagId] = useState("");

  const token = Cookies.get("hwh-jwt");

  useEffect(() => {
    //Check if user is logged in
    if(token === undefined) {
      navigate("/login");
      return;
    }
    if (user === null) {
      dispatch(fetchUserData(token)); 
    }    
    getTags();
  }, []);

  useEffect(() => {
    getThreads(search);    
  },[page, count, filterByTagID]);

  useEffect(() => {
    let newPaginationItems = [];
    for (let num = 1; num < Math.ceil(totalThreads/count)+1; num++) {
      newPaginationItems.push(
        <Pagination.Item key={num} active={num === page} onClick={() => setPage(num)}>
          {num}
        </Pagination.Item>);
    }
    setPaginationItems(newPaginationItems);
  },[totalThreads, page]);

  useEffect(() => {
    if (error) {
      setAlertMessage(error);
      setAlertVisible(true);
    }    
  }, [error]);

  const handleSearchBarSubmit = (event: SyntheticEvent) => {
    event.preventDefault();
    if (search === "") {
      return;
    }
    console.log("Searching for threads with:", search);
    if (page != 1) {
      setPage(1);
    } else {
      getThreads(search); 
    }
    setSearching(true);
  };

  const handleSearchBack = () => {
    setSearching(false);
    setSearch("");
    if (page != 1) {
      setPage(1);
    } else {
      getThreads(""); 
    }
  };

  const getThreads = (searchParams:string) => {
    console.log("Fetching Posts - Page:", page, "Count:", count, "TagID:", filterByTagID, "Search:", searchParams);
    dispatch(fetchThreadData({ token: token, page: page, count: count, tagId: filterByTagID, search: searchParams })); 
  };

  const getTags = () => {
    console.log("Fetching Tags");
    dispatch(fetchTagData(token));
  }

  return (
    <div>
      {alertVisible && alertMessage !== "" && <Alert variant="danger">{alertMessage}</Alert>}

      <div className="mt-5 mb-3">
        <Form className="d-flex" onSubmit={handleSearchBarSubmit}>
          {searching && <Button className="bg-transparent border-0 text-dark" onClick={handleSearchBack}><ArrowBackIcon style={{fontSize:35}} /></Button>}
          <Form.Control className="me-2" type="text" placeholder="Search..." 
            value={search} onChange={e => setSearch(e.target.value)} />
          <Button variant="primary" className="px-3 py-2" type="submit"><SearchIcon/></Button>
        </Form>
      </div>

      <div className="d-flex mt-4 justify-content-between align-items-stretch flex-wrap">
        <TagsDropdown selectTagFn={setFilteredTagId} forFilter />

        <Button variant="warning" className="mb-3" style={{height:"100%"}}><NavLink  style={{ textDecoration: "none", color: '#000' }} to="/createThread">Create New Thread</NavLink></Button>
      </div>
      
      <DisplayThreads navigate={navigate} />

      <div className="mt-4 d-flex justify-content-center">
        <Pagination>{paginationItems}</Pagination>
      </div>
      
    </div>
  );
}
