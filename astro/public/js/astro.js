/** @jsx React.DOM */

var EditableCell = React.createClass({
    getInitialState: function () {
        return {isEditMode: false, data: ""};
    },
    componentWillMount: function () {
        this.setState({isEditMode: this.props.isEditMode, data: this.props.data});
    },
    handleEditCell: function(evt) {
        this.setState({isEditMode: true});

    },
    handleKeyDown: function(evt) {
        switch(evt.keyCode) {
            case 13: // Enter
            case 9:  // Tab
                this.setState({isEditMode: false});

                $.ajax({
                    url: "/save",
                    type: "POST",
                    data: {id: this.props.id, field: this.props.field, value: evt.target.value},
                    success: function(data) {
                        console.log(data);
                    }.bind(this),
                    error: function(xhr, status, err) {
                        console.error(this.props.url, status, err.toString());
                    }.bind(this)
                });

                break;
        }
    },
    handleChange: function(evt) {
        this.setState({data: evt.target.value});
    },
    render: function() {
        var cellHtml;
        if (this.state.isEditMode) {
            cellHtml = <input type='text' value={this.state.data} onKeyDown={this.handleKeyDown} onChange={this.handleChange} />
        }
        else {
            cellHtml = <div onClick={this.handleEditCell}>{this.state.data}</div>
        }
        return (
        <td>{cellHtml}</td>
            );
    }
});

var Movie = React.createClass({
    render: function() {
        return (
            <tr>
                <EditableCell id={this.props.id} field="title" data={this.props.title} />
                <EditableCell id={this.props.id} field="rank" data={this.props.rank} className="har" />
                <EditableCell id={this.props.id} field="year" data={this.props.year} />
                <EditableCell id={this.props.id} field="rating" data={this.props.rating} />
                <EditableCell id={this.props.id} field="reviews" data={this.props.reviews} />
            </tr>
            );
    }
});

var MovieList = React.createClass({
    getInitialState: function() {
        return {data: []};
    },
    componentWillMount: function() {
        $.ajax({
            url: this.props.url,
            dataType: 'json',
            success: function(data) {
                this.setState({data: data});
            }.bind(this),
            error: function(xhr, status, err) {
                console.error(this.props.url, status, err.toString());
            }.bind(this)
        });
    },
   render: function() {
       var movies = this.state.data.map(function (movie) {
           return <Movie id={movie.id} title={movie.title} rank={movie.rank} year={movie.year} rating={movie.rating} reviews={movie.reviews} />;
       });
       return (
           <tbody>{movies}</tbody>
           );
   }
});

React.renderComponent(
    <table id="movie-table" className="tablesaw" data-mode="stack" data-sortable>
        <thead>
            <tr>
                <th data-sortable-col data-sortable-default-col data-priority="persist">Movie Title</th>
                <th data-sortable-col data-priority="3">Rank</th>
                <th data-sortable-col data-priority="2">Year</th>
                <th data-sortable-col data-priority="1">
                    <abbr title="Rotten Tomato Rating">Rating</abbr>
                </th>
                <th data-sortable-col data-priority="4">Reviews</th>
            </tr>
        </thead>
        <MovieList url="/public/js/movies.json"/>
    </table>,
    document.getElementById('table-container')
);

$(document).trigger("enhance.tablesaw");