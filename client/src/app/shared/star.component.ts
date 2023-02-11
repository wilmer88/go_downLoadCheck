import {Component,EventEmitter,Input, OnChanges, Output} from "@angular/core";


@Component({
    selector: "happy-stars",
    templateUrl: "./star.component.html",
    styleUrls:["./star.component.css"]
})


export class StarComponent{
   @Input() rating: number = 0;
    cropWidth: number = 75;
   @Output() ratingClicked: EventEmitter<string>=
    new EventEmitter
    
    ngOnChanges(): void {
        this.cropWidth= this.rating * 75/5;
    }

    onClick(): void {
        this.ratingClicked.emit(`the ${this.rating} star rating was clicked`)

    }

};