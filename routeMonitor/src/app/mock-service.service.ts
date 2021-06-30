import { Injectable } from '@angular/core';
import mockData from './routes/routes.json'
import {webSocket} from 'rxjs/webSocket'
import { Observable } from 'rxjs';

const connection = webSocket('ws://localhost:8080')

export interface Route {
  id: number;
  driver_name: string;
  created_at: string;
  completed_at: string;
  deliveries: number;
  status: string;
}

export interface RootObject {
  routes: Route[];
}

@Injectable({
  providedIn: 'root'
})
export class MockServiceService {

  public Rutas:any = []

  constructor() { }

  public GetSockete(): Observable<any>{
    
    return connection
  }

  public GetMock(){

    this.Rutas = mockData.routes

    return this.Rutas
    
  }
  

  
}
