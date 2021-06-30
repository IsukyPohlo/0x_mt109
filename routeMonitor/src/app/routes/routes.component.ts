import { Component, OnInit } from '@angular/core';
import { MockServiceService }from '../mock-service.service'

@Component({
  selector: 'app-routes',
  templateUrl: './routes.component.html',
  styleUrls: ['./routes.component.css']
})

export class RoutesComponent implements OnInit {

  public data:any = []

  constructor(public serv:MockServiceService) { }

  ngOnInit(): void {

    this.data = this.serv.GetMock()

    //console.log(this.data.routes)

    this.serv.GetSockete().subscribe((msg:any)=>{

      //Actualizar registro
      let indChange = this.data.findIndex(((obj:any)=>obj.id==msg.route_id))
    
      this.data[indChange].completed_at = msg.completed_at
      this.data[indChange].deliveries = msg.deliveries
      this.data[indChange].status = msg.status

    },error=>{
     //Poserror
    },()=>{
     //Se detiene
    })

  }

}
