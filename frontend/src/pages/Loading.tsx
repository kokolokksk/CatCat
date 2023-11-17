
import React from 'react';
import { useNavigate } from 'react-router-dom';

interface LoadingProps {
    path: string;
}

const Loading: React.FC<LoadingProps> = (_props: LoadingProps) => {
    const { path } = _props;
    let temp = path;
    console.log(path);
    if(path === "/") {
        temp = "/setting?loading"
    }else{
        temp =  temp.replace("?loading", "");
    }
    // pealse go the path by nagivat.
    const navigate = useNavigate();
    setTimeout(() => {
        navigate(temp);
    }, 10000);
    return (
        <div className="loading-container">
            <div className="loading-spinner"></div>
        </div>
    );
};

export default Loading;
